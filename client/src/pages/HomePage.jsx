import React, { useState } from 'react'
import LoginPage from './LoginPage'
import RegisterPage from './RegisterPage'

function Home() {
    const [view, setView] = useState(null); // null = chưa chọn

    return (
        <div>
            <h1>Welcome to TDOJ</h1>
            <div>
                <button onClick={() => setView("login")}>Đăng nhập</button>
                <button onClick={() => setView("register")}>Đăng ký</button>
            </div>
            <div>
                {view === "login" && <LoginPage />}
                {view === "register" && <RegisterPage />}
            </div>
        </div>
    )
}

export default Home
