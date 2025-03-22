from pydantic import BaseModel, ConfigDict
from typing import Optional
from datetime import datetime

class ReservationBase(BaseModel):
    user_id: str
    barber_id: str
    appointment_time: datetime

class ReservationCreate(ReservationBase):
    pass

class ReservationUpdate(BaseModel):
    appointment_time: Optional[datetime] = None
    status: Optional[str] = None

class Reservation(ReservationBase):
    id: str
    status: str = "pending"

    model_config = ConfigDict(from_attributes=True)
