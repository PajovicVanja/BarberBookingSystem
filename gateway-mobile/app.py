from flask import Flask, request, Response, jsonify
import requests

app = Flask(__name__)

def forward_request(target_url, subpath):
    url = f"{target_url}{subpath}"
    if request.query_string:
        url = f"{url}?{request.query_string.decode()}"
    
    headers = {key: value for key, value in request.headers if key.lower() != 'host'}
    
    resp = requests.request(
        method=request.method,
        url=url,
        headers=headers,
        data=request.get_data(),
        cookies=request.cookies,
        allow_redirects=False
    )
    
    excluded_headers = ['content-encoding', 'content-length', 'transfer-encoding', 'connection']
    response_headers = [(name, value) for name, value in resp.raw.headers.items() if name.lower() not in excluded_headers]
    return Response(resp.content, resp.status_code, response_headers)

# Proxy for Payment Service
@app.route('/m/payments/<path:subpath>', methods=["GET", "POST", "PUT", "DELETE", "PATCH"])
def proxy_payments(subpath):
    target = "http://payment-service:8080/api/payments/"
    return forward_request(target, subpath)

# Proxy for Reservation Service
@app.route('/m/reservations/<path:subpath>', methods=["GET", "POST", "PUT", "DELETE", "PATCH"])
def proxy_reservations(subpath):
    target = "http://reservation-service:8000/api/reservations/"
    return forward_request(target, subpath)

# Proxy for User Service
@app.route('/m/users/<path:subpath>', methods=["GET", "POST", "PUT", "DELETE", "PATCH"])
def proxy_users(subpath):
    target = "http://user-service:3000/api/users/"
    return forward_request(target, subpath)

# 🔥 NEW aggregated route for mobile (similar to web gateway)
@app.route('/m/dashboard/user/<id>', methods=["GET"])
def dashboard_user(id):
    try:
        headers = {key: value for key, value in request.headers if key.lower() != 'host'}
        
        # concurrent fetching using simple requests (not fully async but OK for now)
        profile_resp = requests.get(f"http://user-service:3000/api/users/profile", headers=headers, timeout=2)
        reservations_resp = requests.get(f"http://reservation-service:8000/api/reservations/user/{id}", timeout=2)
        payments_resp = requests.get(f"http://payment-service:8080/api/payments/user/{id}", timeout=2)
        
        # Build JSON response
        profile = profile_resp.json()
        reservations = reservations_resp.json()
        payments = payments_resp.json()

        return jsonify({
            "profile": profile,
            "reservations": reservations,
            "payments": payments
        })

    except Exception as e:
        print(f"Aggregation failed: {str(e)}")
        return jsonify({"error": "aggregation-failed", "detail": str(e)}), 502

# Root endpoint
@app.route('/')
def home():
    return {"message": "Welcome to the Mobile API Gateway (Python)"}

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
