"use client"

import { useState } from "react"
import axios from "axios"
import { useRouter } from "next/navigation"
import { loginUser, registerUser } from "../service/user"

export default function LoginPage() {
  const router = useRouter()

  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")

  const handleRegister= async () => {
    try {
      setLoading(true)
      setError("")
      let { data } = await registerUser({
        email,
        password,
      })

      // ✅ เก็บ token
      localStorage.setItem("token", data.token)

      // ✅ redirect
      router.push("/")

    } catch (err: any) {
      setError(err.response?.data?.error || "login failed")
    } finally {
      setLoading(false)
    }
  }

    const handleLogin= async () => {
    try {
      setLoading(true)
      setError("")
      let { data } = await loginUser({
        email,
        password,
      })

      // ✅ เก็บ token
      localStorage.setItem("token", data.token)

      // ✅ redirect
      router.push("/")

    } catch (err: any) {
      setError(err.response?.data?.error || "login failed")
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="flex items-center justify-center h-screen">
      <div className="w-80 p-6 border rounded">
        <h1 className="text-xl mb-4">Login</h1>

        {error && <p className="text-red-500">{error}</p>}

        <input
          type="email"
          placeholder="Email"
          className="w-full mb-2 p-2 border"
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          placeholder="Password"
          className="w-full mb-4 p-2 border"
          onChange={(e) => setPassword(e.target.value)}
        />

        <button
          onClick={handleLogin}
          disabled={loading}
          className="w-full bg-blue-500 text-white p-2"
        >
          {loading ? "Loading..." : "Login"}
        </button>
      </div>
    </div>
  )
}