from flask import Flask, request, Response
import requests

app = Flask(__name__)

def forward_request(target_url, subpath):
    """
    Forwards the incoming request to the target URL plus the subpath,
    preserving method, headers, query parameters, and body.
    """
    # Construct full URL including query parameters.
    url = f"{target_url}{subpath}"
    if request.query_string:
        url = f"{url}?{request.query_string.decode()}"
    
    # Prepare headers (exclude host header to let requests set it properly).
    headers = {key: value for key, value in request.headers if key.lower() != 'host'}
    
    # Make the request to the target service.
    resp = requests.request(
        method=request.method,
        url=url,
        headers=headers,
        data=request.get_data(),
        cookies=request.cookies,
        allow_redirects=False
    )
    
    # Exclude certain headers and pass the response back.
    excluded_headers = ['content-encoding', 'content-length', 'transfer-encoding', 'connection']
    response_headers = [(name, value) for name, value in resp.raw.headers.items() if name.lower() not in excluded_headers]
    return Response(resp.content, resp.status_code, response_headers)

# Proxy for Payment Service (assumed to be available at port 8080)
@app.route('/m/payments/<path:subpath>', methods=["GET", "POST", "PUT", "DELETE", "PATCH"])
def proxy_payments(subpath):
    target = "http://payment-service:8080/api/payments/"
    return forward_request(target, subpath)

# Proxy for Reservation Service (assumed to be available at port 8000)
@app.route('/m/reservations/<path:subpath>', methods=["GET", "POST", "PUT", "DELETE", "PATCH"])
def proxy_reservations(subpath):
    target = "http://reservation-service:8000/api/reservations/"
    return forward_request(target, subpath)

# Proxy for User Service (assumed to be available at port 3000)
@app.route('/m/users/<path:subpath>', methods=["GET", "POST", "PUT", "DELETE", "PATCH"])
def proxy_users(subpath):
    target = "http://user-service:3000/api/users/"
    return forward_request(target, subpath)

# A simple root endpoint for health checking.
@app.route('/')
def home():
    return {"message": "Welcome to the Mobile API Gateway (Python)"}

if __name__ == '__main__':
    # Run on all network interfaces so Docker can expose the port.
    app.run(host='0.0.0.0', port=5000)
