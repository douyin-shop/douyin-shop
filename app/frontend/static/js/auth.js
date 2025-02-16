// const API_BASE = 'http://127.0.0.1:8080';
const API_BASE = '';
function showForm(formType) {
    document.getElementById('loginForm').classList.toggle('hidden', formType !== 'login');
    document.getElementById('registerForm').classList.toggle('hidden', formType !== 'register');
}

async function handleLogin(e) {
    e.preventDefault();
    const email = document.getElementById('loginEmail').value;
    const password = document.getElementById('loginPassword').value;

    try {
        const response = await fetch(`${API_BASE}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });


        // 正确的返回结果
        `{
            "code": 0,
            "data": {
                "user_id": 2,
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzgzODg0MTMsInVzZXJfaWQiOjJ9.PoH-e_v19H3Zk1IAbHeW0V7hKDc7y5DnyCcmLFB_RdU"
            },
            "msg": "success"
        }`
        const data = await response.json();

        // 判断code是否为0
        if (data.code === 0) {
            localStorage.setItem('authToken', data.data.token);
            window.location.href = 'products.html';
        }else{
            alert('Login failed: ' + (data.msg || 'Unknown error'));
        }

    } catch (error) {
        console.error('Login error:', error);
        alert('Login failed. Please try again.');
    }
}

async function handleRegister(e) {
    e.preventDefault();
    const email = document.getElementById('registerEmail').value;
    const password = document.getElementById('registerPassword').value;

    try {
        const response = await fetch(`${API_BASE}/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password })
        });

        const data = await response.json();
        if (data.code === 0) {
            alert('Registration successful! Please login.');
            showForm('login');
        } else {
            alert('Registration failed: ' + (data.msg || 'Unknown error'));
        }

    } catch (error) {
        console.error('Registration error:', error);
        alert('Registration failed. Please try again.');
    }
}