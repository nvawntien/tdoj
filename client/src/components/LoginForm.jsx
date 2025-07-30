import React, { useState } from 'react'
import { loginUser } from "../services/authService"

function LoginForm() {
    const [formData, setFormData] = useState({
        email: '',
        password: '',
    })

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        })
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const data = await loginUser(formData);
            alert("Đăng nhập thành công!");
            console.log("Server response:", data);
        } catch (err) {
            console.error("Lỗi đăng nhập:", err);
            alert("Đăng nhập thất bại!");
            console.error("Full error:", err);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Email:
                <input type="email" name="email" value={formData.email} onChange={handleChange} required />
            </label>
            <br />

            <label>
                Mật khẩu:
                <input type="password" name="password" value={formData.password} onChange={handleChange} required />
            </label>
            <br />

            <button type="submit">Đăng nhập</button>
        </form>
    )
}

export default LoginForm


    