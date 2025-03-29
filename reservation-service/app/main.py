from fastapi import FastAPI, HTTPException
from app.models import Reservation, ReservationCreate, ReservationUpdate
from app import crud
from app.utils.rabbitmq_consumer import run_consumer_in_background
import uvicorn
from app.utils.logger import logger
from bson.errors import InvalidId  # ⬅️ Add this import
from typing import List
import logging
logging.basicConfig(level=logging.DEBUG)  # 👈 enable DEBUG level globally


app = FastAPI(title="Reservation Service")

@app.on_event("startup")
async def startup_event():
    # Start the RabbitMQ consumer in a background thread.
    run_consumer_in_background()

@app.post("/api/reservations", response_model=Reservation)
async def create_reservation(reservation: ReservationCreate):
    logger.info("Creating reservation")
    created = await crud.create_reservation(reservation)
    return created

@app.get("/api/reservations/{reservation_id}", response_model=Reservation)
async def get_reservation(reservation_id: str):
    logger.info(f"Fetching reservation {reservation_id}")
    try:
        reservation = await crud.get_reservation(reservation_id)
    except InvalidId:
        raise HTTPException(status_code=400, detail="Invalid reservation ID format")
    if not reservation:
        raise HTTPException(status_code=404, detail="Reservation not found")
    return reservation

@app.get("/api/reservations/user/{user_id}", response_model=List[Reservation])
async def list_user_reservations(user_id: str):
    logger.info(f"🔍 Fetching reservations for user: {user_id}")
    reservations = await crud.list_user_reservations(user_id)
    logger.info(f"Found {len(reservations)} reservations")
    return reservations


@app.patch("/api/reservations/{reservation_id}", response_model=Reservation)
async def update_reservation(reservation_id: str, reservation_update: ReservationUpdate):
    logger.info(f"Updating reservation {reservation_id} with {reservation_update}")
    updated = await crud.update_reservation(reservation_id, reservation_update)
    if not updated:
        raise HTTPException(status_code=404, detail="Reservation not found")
    return updated

@app.delete("/api/reservations/{reservation_id}")
async def delete_reservation(reservation_id: str):
    logger.info(f"Deleting reservation {reservation_id}")
    success = await crud.delete_reservation(reservation_id)
    if not success:
        raise HTTPException(status_code=404, detail="Reservation not found")
    return {"detail": "Reservation deleted successfully"}

@app.post("/api/reservations/confirm", response_model=Reservation)
async def confirm_reservation(reservation: ReservationCreate):
    logger.info("Confirming reservation payment")
    created = await crud.create_reservation(reservation)
    # Publish confirmation message to RabbitMQ
    from app.utils.rabbitmq import send_confirmation
    send_confirmation(created.id)
    # Update status to confirmed
    from app.models import ReservationUpdate
    update = ReservationUpdate(status="confirmed")
    updated = await crud.update_reservation(created.id, update)
    return updated

if __name__ == "__main__":
    uvicorn.run("app.main:app", host="0.0.0.0", port=8000, reload=True)
