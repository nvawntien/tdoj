// src/services/authService.js
import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080'; // chỉnh lại nếu BE khác cổng

export async function registerUser(userData) {
  try {
    const res = await axios.post(`${API_BASE_URL}/tdoj/user/register`, userData);
    return res.data;
  } catch (err) {
    throw err.response?.data?.message || 'Đăng ký thất bại';
  }
}
