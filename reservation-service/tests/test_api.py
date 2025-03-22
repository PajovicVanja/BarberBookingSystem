import httpx
from datetime import datetime
import pytest

BASE_URL = "http://127.0.0.1:8000/api/reservations"

@pytest.mark.asyncio
async def test_reservation_workflow():
    async with httpx.AsyncClient(timeout=10.0) as client:
        print("📦 Creating reservation...")
        reservation_data = {
            "user_id": "user123",
            "barber_id": "barber456",
            "appointment_time": datetime.now().isoformat()
        }
        response = await client.post(BASE_URL, json=reservation_data)
        assert response.status_code == 200
        created = response.json()
        reservation_id = created["id"]
        print("✅ Reservation created:", created)

        print("🔍 Getting reservation...")
        response = await client.get(f"{BASE_URL}/{reservation_id}")
        assert response.status_code == 200
        assert response.json()["id"] == reservation_id
        print("✅ Reservation fetched:", response.json())

        print("✏️ Updating reservation...")
        updated_data = {
            "appointment_time": datetime.now().isoformat(),
            "status": "confirmed"
        }
        response = await client.patch(f"{BASE_URL}/{reservation_id}", json=updated_data)
        assert response.status_code == 200
        print("✅ Reservation updated:", response.json())

        print("📋 Listing reservations for user...")
        response = await client.get(f"{BASE_URL}/user/{reservation_data['user_id']}")
        assert response.status_code == 200
        print("✅ Reservations found:", len(response.json()))

        print("🗑️ Deleting reservation...")
        response = await client.delete(f"{BASE_URL}/{reservation_id}")
        assert response.status_code == 200
        print("✅ Reservation deleted.")
