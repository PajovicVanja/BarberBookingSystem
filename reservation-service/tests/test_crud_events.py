# tests/test_crud_events.py
import pytest
from datetime import datetime
from bson import ObjectId
from app.crud import create_reservation
from app.models import ReservationCreate
from unittest.mock import patch, MagicMock

@pytest.mark.asyncio
async def test_create_reservation_publishes_event(monkeypatch):
    # 1) Stub out the DB insert so we control the inserted_id
    fake_insert = MagicMock()
    fake_insert.inserted_id = ObjectId("656033f1f1a2f3e4d5c6b7a8")
    # Make insert_one an async function that returns our fake
    async def fake_insert_one(doc):
        return fake_insert
    # Patch the actual collection.insert_one method
    import app.crud as crud_module
    monkeypatch.setattr(crud_module.collection, "insert_one", fake_insert_one)
    # 2) Capture publish_event calls
    published = {}
    def fake_publish(event_type, data):
        published['type'] = event_type
        published['data'] = data
    monkeypatch.setattr("app.crud.publish_event", fake_publish)

    # 3) Call create_reservation
    dto = ReservationCreate(
        user_id="u1", barber_id="b2", appointment_time=datetime(2025,5,18,10,0,0)
    )
    res = await create_reservation(dto)

    # 4) Assertions
    assert res.id == "656033f1f1a2f3e4d5c6b7a8"
    assert published['type'] == "ReservationCreated"
    assert published['data']["user_id"] == "u1"
    assert published['data']["barber_id"] == "b2"
    assert published['data']["appointment_time"] == "2025-05-18T10:00:00"
    assert published['data']["status"] == "pending"
