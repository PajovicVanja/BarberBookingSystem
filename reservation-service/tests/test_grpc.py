import sys
import os
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

import pytest
import grpc
from datetime import datetime

from app import reservation_pb2, reservation_pb2_grpc

@pytest.fixture(scope="module")
def grpc_stub():
    channel = grpc.insecure_channel("localhost:50051")
    stub = reservation_pb2_grpc.ReservationServiceStub(channel)
    return stub

def test_grpc_reservation_flow(grpc_stub):
    print("📦 gRPC: Creating reservation...")
    create_response = grpc_stub.CreateReservation(
        reservation_pb2.ReservationRequest(
            user_id="user123",
            barber_id="barber456",
            datetime=datetime.now().isoformat()
        )
    )
    assert create_response.success
    reservation = create_response.reservation
    reservation_id = reservation.id
    print("✅ Created:", reservation)

    print("🔍 gRPC: Getting reservation...")
    get_response = grpc_stub.GetReservation(
        reservation_pb2.GetReservationRequest(id=reservation_id)
    )
    assert get_response.success
    assert get_response.reservation.id == reservation_id
    print("✅ Fetched:", get_response.reservation)

    print("✏️ gRPC: Updating reservation...")
    update_response = grpc_stub.UpdateReservation(
        reservation_pb2.UpdateReservationRequest(
            id=reservation_id,
            datetime=datetime.now().isoformat()
        )
    )
    assert update_response.success
    print("✅ Updated:", update_response.reservation)

    print("📋 gRPC: Listing reservations...")
    list_response = grpc_stub.ListUserReservations(
        reservation_pb2.ListUserReservationsRequest(user_id="user123")
    )
    assert len(list_response.reservations) > 0
    assert any(r.id == reservation_id for r in list_response.reservations)
    print(f"✅ Found {len(list_response.reservations)} reservations")

    print("💳 gRPC: Confirming payment...")
    confirm_response = grpc_stub.ConfirmPayment(
        reservation_pb2.ConfirmPaymentRequest(id=reservation_id)
    )
    assert confirm_response.success
    assert confirm_response.reservation.status == "confirmed"
    print("✅ Payment confirmed")

    print("🗑️ gRPC: Cancelling reservation...")
    cancel_response = grpc_stub.CancelReservation(
        reservation_pb2.CancelReservationRequest(id=reservation_id)
    )
    assert cancel_response.success
    print("✅ Reservation cancelled")
