from concurrent import futures
import grpc
import time
import asyncio

from app.config import GRPC_SERVER_PORT
from app import crud
from app.models import ReservationCreate, ReservationUpdate
from app.utils.logger import logger
from app import reservation_pb2
from app import reservation_pb2_grpc


class ReservationServiceServicer(reservation_pb2_grpc.ReservationServiceServicer):
    def __init__(self, loop):
        self.loop = loop
        print("✅ ReservationServiceServicer initialized")

    def CreateReservation(self, request, context):
        print("📦 CreateReservation RPC called")
        reservation_data = {
            "user_id": request.user_id,
            "barber_id": request.barber_id,
            "appointment_time": request.datetime,
        }
        reservation_create = ReservationCreate(**reservation_data)
        reservation = self.loop.run_until_complete(crud.create_reservation(reservation_create))
        return reservation_pb2.ReservationResponse(
            success=True,
            message="Reservation created",
            reservation=reservation_pb2.Reservation(
                id=reservation.id,
                user_id=reservation.user_id,
                barber_id=reservation.barber_id,
                datetime=str(reservation.appointment_time),
                status=reservation.status,
            )
        )

    def GetReservation(self, request, context):
        print("🔍 GetReservation RPC called")
        reservation = self.loop.run_until_complete(crud.get_reservation(request.id))
        if reservation:
            return reservation_pb2.ReservationResponse(
                success=True,
                message="Reservation found",
                reservation=reservation_pb2.Reservation(
                    id=reservation.id,
                    user_id=reservation.user_id,
                    barber_id=reservation.barber_id,
                    datetime=str(reservation.appointment_time),
                    status=reservation.status,
                )
            )
        return reservation_pb2.ReservationResponse(
            success=False,
            message="Reservation not found"
        )

    def CancelReservation(self, request, context):
        print("🗑️ CancelReservation RPC called")
        success = self.loop.run_until_complete(crud.delete_reservation(request.id))
        return reservation_pb2.ReservationResponse(
            success=success,
            message="Reservation cancelled" if success else "Reservation not found"
        )

    def UpdateReservation(self, request, context):
        print("✏️ UpdateReservation RPC called")
        update_data = {}
        if request.datetime:
            update_data["appointment_time"] = request.datetime
        reservation_update = ReservationUpdate(**update_data)
        reservation = self.loop.run_until_complete(
            crud.update_reservation(request.id, reservation_update)
        )
        return reservation_pb2.ReservationResponse(
            success=True,
            message="Reservation updated",
            reservation=reservation_pb2.Reservation(
                id=reservation.id,
                user_id=reservation.user_id,
                barber_id=reservation.barber_id,
                datetime=str(reservation.appointment_time),
                status=reservation.status,
            )
        )

    def ListUserReservations(self, request, context):
        print("📋 ListUserReservations RPC called")
        reservations = self.loop.run_until_complete(crud.list_user_reservations(request.user_id))
        response = reservation_pb2.ReservationListResponse()
        for res in reservations:
            reservation_proto = reservation_pb2.Reservation(
                id=res.id,
                user_id=res.user_id,
                barber_id=res.barber_id,
                datetime=str(res.appointment_time),
                status=res.status,
            )
            response.reservations.append(reservation_proto)
        return response

    def ConfirmPayment(self, request, context):
        print("💳 ConfirmPayment RPC called")
        update = ReservationUpdate(status="confirmed")
        reservation = self.loop.run_until_complete(crud.update_reservation(request.id, update))
        return reservation_pb2.ReservationResponse(
            success=True,
            message="Payment confirmed and reservation finalized",
            reservation=reservation_pb2.Reservation(
                id=reservation.id,
                user_id=reservation.user_id,
                barber_id=reservation.barber_id,
                datetime=str(reservation.appointment_time),
                status=reservation.status,
            )
        )


def serve():
    print("🚀 Starting gRPC server setup...")
    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    reservation_pb2_grpc.add_ReservationServiceServicer_to_server(
        ReservationServiceServicer(loop), server
    )
    server.add_insecure_port(f"[::]:{GRPC_SERVER_PORT}")
    server.start()
    print(f"✅ gRPC server started on port {GRPC_SERVER_PORT}")
    logger.info(f"gRPC server started on port {GRPC_SERVER_PORT}")
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        print("🛑 gRPC server shutting down")
        server.stop(0)


if __name__ == "__main__":
    print("📦 grpc_server.py __main__ block running...")
    serve()
