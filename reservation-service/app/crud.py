from app.config import MONGO_URL, DATABASE_NAME
import logging
from typing import List
from app.models import Reservation, ReservationCreate, ReservationUpdate
from motor.motor_asyncio import AsyncIOMotorClient
from bson import ObjectId

logger = logging.getLogger(__name__)
client = AsyncIOMotorClient(MONGO_URL)
db = client[DATABASE_NAME]
collection = db["reservations"]

async def create_reservation(reservation: ReservationCreate) -> Reservation:
    reservation_dict = reservation.model_dump()
    reservation_dict["status"] = "pending"
    # note: message will be absent until a decline
    result = await collection.insert_one(reservation_dict)
    reservation_dict["id"] = str(result.inserted_id)
    return Reservation(**reservation_dict)

async def get_reservation(reservation_id: str) -> Reservation:
    document = await collection.find_one({"_id": ObjectId(reservation_id)})
    if document:
        document["id"] = str(document["_id"])
        return Reservation(**document)

async def update_reservation(reservation_id: str, reservation_update: ReservationUpdate) -> Reservation:
    logger.info(f"Patch received: {reservation_update}")
    update_data = {k: v for k, v in reservation_update.model_dump().items() if v is not None}
    logger.info(f"Updating DB with: {update_data}")
    await collection.update_one({"_id": ObjectId(reservation_id)}, {"$set": update_data})
    return await get_reservation(reservation_id)

async def delete_reservation(reservation_id: str) -> bool:
    result = await collection.delete_one({"_id": ObjectId(reservation_id)})
    return result.deleted_count > 0

async def list_user_reservations(user_id: str) -> List[Reservation]:
    logger.info("🔍 Fetching reservations for user: %s", user_id)
    results = await collection.find({"user_id": user_id}).to_list(length=100)
    reservations: List[Reservation] = []
    for i, res in enumerate(results):
        try:
            reservations.append(Reservation(
                id=str(res["_id"]),
                user_id=res["user_id"],
                barber_id=res["barber_id"],
                appointment_time=res["appointment_time"],
                status=res.get("status", "pending"),
                message=res.get("message")
            ))
        except Exception as e:
            logger.error("❌ Skipping malformed reservation: %s", res)
            logger.error("📛 Reason: %s", e)
    return reservations

async def list_barber_reservations(barber_id: str) -> list[Reservation]:
    results = await collection.find({"barber_id": barber_id}).to_list(length=100)
    reservations: list[Reservation] = []
    for res in results:
        reservations.append(Reservation(
            id=str(res["_id"]),
            user_id=res["user_id"],
            barber_id=res["barber_id"],
            appointment_time=res["appointment_time"],
            status=res.get("status", "pending"),
            message=res.get("message")
        ))
    return reservations
