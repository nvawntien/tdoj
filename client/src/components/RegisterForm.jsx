import React, { useState } from 'react'
import { registerUser } from "../services/authService"

function RegisterForm() {
    const [formData, setFormData] = useState({
        full_name: '',
        email: '',
        username: '',
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
            const data = await registerUser(formData);
            alert("Đăng ký thành công!");
            console.log("Server response:", data);
        } catch (err) {
            console.error("Lỗi đăng ký:", err);
            alert("Đăng ký thất bại!");
            console.error("Full error:", err);
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Họ và tên:
                    <input type="text" name="full_name" value={formData.full_name} onChange={handleChange} required />
            </label>
            <br />

            <label>
                Email:
                    <input type="email" name="email" value={formData.email} onChange={handleChange} required />
            </label>
            <br />
        
            <label>
                Tên đăng nhập:
                    <input type="text" name="username" value={formData.username} onChange={handleChange} required />
            </label>
            <br />
            
            <label>
                Mật khẩu:
                    <input type="password" name="password" value={formData.password} onChange={handleChange} required />
            </label>
            <br />

        <button type="submit">Đăng ký</button>
        </form>
    )
    }

export default RegisterForm
