import axios from 'axios'
const SERVER_URI = "http://localhost:4000"

export async function getAll() {
    return await axios.get(SERVER_URI + "/users")
}

export async function create(payload) {
    return await axios.post(SERVER_URI + "/users", payload)
}