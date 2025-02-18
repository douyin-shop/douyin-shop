const API_BASE = '';

function showForm(formType) {
    const container = document.querySelector('.form-container');
    document.getElementById('loginForm').classList.toggle('hidden', formType !== 'login');
    document.getElementById('registerForm').classList.toggle('hidden', formType !== 'register');

    // 更新切换动画
    if(formType === 'register') {
        container.querySelector('.form-toggle').classList.add('register-active');
    } else {
        container.querySelector('.form-toggle').classList.remove('register-active');
    }
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

        const data = await response.json();
        if (data.code === 0) {
            localStorage.setItem('authToken', data.data.token);
            window.location.href = 'products.html';
        } else {
            alert('登录失败: ' + (data.msg || '请检查输入信息'));
        }
    } catch (error) {
        console.error('登录错误:', error);
        alert('登录失败，请稍后重试');
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
            alert('注册成功，请登录');
            showForm('login');
            document.getElementById('loginEmail').value = email;
        } else {
            alert('注册失败: ' + (data.msg || '该邮箱已被使用'));
        }
    } catch (error) {
        console.error('注册错误:', error);
        alert('注册失败，请检查网络连接');
    }
}