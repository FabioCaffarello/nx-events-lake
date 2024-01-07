import os
import requests
from datetime import datetime, timedelta
from flask import Flask, redirect, request, jsonify, session
import urllib.parse

app = Flask(__name__)
app.secret_key = os.getenv("FLASK_SPOTIFY_INTEGRATION_SECRET_KEY")


CLIENT_ID = os.getenv("SPOTIFY_CLIENT_ID")
CLIENT_SECRET = os.getenv("SPOTIFY_CLIENT_SECRET")
REDIRECT_URI = os.getenv("SPOTIFY_REDIRECT_URI")

ENVIRONMENT = os.getenv("ENVIRONMENT", "development")


AUTH_URL = "https://accounts.spotify.com/authorize"
TOKEN_URL = "https://accounts.spotify.com/api/token"
API_BASE_URL = "https://api.spotify.com/v1"


@app.route("/")
def index():
    return "Welcome to Spotify Integration App <a href='/login'> Login with Spotify</a>"


@app.route("/login")
def login():
    scope = "user-read-email user-read-private"

    params = {
        "client_id": CLIENT_ID,
        "response_type": "code",
        "redirect_uri": REDIRECT_URI,
        "scope": scope,
        "show_dialog": True if ENVIRONMENT == "development" else False, # show dialog in development mode only for debugging
    }

    auth_url = f"{AUTH_URL}?{urllib.parse.urlencode(params)}"

    return redirect(auth_url)


@app.route("/callback")
def callback():
    if "error" in request.args:
        return jsonify({"error": request.args["error"]})

    if "code" in request.args:
        req_body = {
            "code": request.args["code"],
            "grant_type": "authorization_code",
            "redirect_uri": REDIRECT_URI,
            "client_id": CLIENT_ID,
            "client_secret": CLIENT_SECRET,
        }

        response = requests.post(TOKEN_URL, data=req_body)
        token_info = response.json()

        session["access_token"] = token_info["access_token"]
        session["refresh_token"] = token_info["refresh_token"]
        session["expires_at"] = datetime.now().timestamp() + token_info["expires_in"]

        return redirect("/playlists")


@app.route("/playlists")
def get_playlists():
    if "access_token" not in session:
        return redirect("/login")

    if session["expires_at"] < datetime.now().timestamp():
        print("Token expired, refreshing...")
        return redirect("/refresh-token")

    headers = {
        "Authorization": f"Bearer {session['access_token']}"
    }
    response = requests.get(f"{API_BASE_URL}/me/playlists", headers=headers)
    playlists = response.json()

    return jsonify(playlists)


@app.route("/refresh-token")
def refresh_token():
    if "refresh_token" not in session:
        return redirect("/login")

    if session["expires_at"] < datetime.now().timestamp():
        print("Token expired, refreshing...")
        req_body = {
            "grant_type": "refresh_token",
            "refresh_token": session["refresh_token"],
            "client_id": CLIENT_ID,
            "client_secret": CLIENT_SECRET,
        }

        response = requests.post(TOKEN_URL, data=req_body)
        new_token_info = response.json()

        session["access_token"] = new_token_info["access_token"]
        session["expires_at"] = datetime.now().timestamp() + new_token_info["expires_in"]

        return redirect("/playlists")


if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=5001)
