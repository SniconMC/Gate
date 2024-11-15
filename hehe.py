import requests
import os
import socket
import time

def register_server():
    server_name = "hub-10"
    server_address = "127.0.0.1:25569"
    auth_token = "yJNabF741tgsBdZbOtOnyloqhCP0O2LCwR4WUVNGkms="
    
    if not server_name or not server_address:
        print("Error: SERVER_NAME or SERVER_ADDRESS is not set!")
        return

    if not auth_token:
        print("Error: GATE_API_AUTH_TOKEN is not set!")
        return

    # Retry loop for DNS resolution
    max_retries = 5
    retry_delay = 10  # seconds

    for attempt in range(max_retries):
        try:
            # Attempt DNS resolution
            socket.gethostbyname(server_address.split(':')[0])
            print(f"DNS resolved successfully: {server_address}")
            break
        except socket.gaierror:
            print(f"DNS resolution failed for {server_address}, attempt {attempt + 1} of {max_retries}")
            if attempt < max_retries - 1:
                time.sleep(retry_delay)
            else:
                print("Max retries reached. Could not resolve DNS.")
                return

    # Proceed with server registration
    gate_api_host = os.getenv('GATE_API_HOST', 'localhost')
    gate_api_port = os.getenv('GATE_API_PORT', '8080')
    url = f"http://{gate_api_host}:{gate_api_port}/addserver"

    headers = {
        "Authorization": auth_token,
        "Content-Type": "application/json"
    }
    payload = {
        "name": server_name,
        "address": server_address,
        "fallback": True  # Modify as needed
    }

    response = requests.post(url, headers=headers, json=payload)
    if response.status_code == 200:
        print(f"Successfully registered server {server_name}.")
    else:
        print(f"Failed to register server {server_name}: {response.content}")

# Example usage
if __name__ == "__main__":
    register_server()
