let token = '';

async function handleLogin() {
  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;
  try {
    const resp = await fetch('http://127.0.0.1:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ email, password }),
    });
    const result = await resp.json();
    if (result.code !== 0) {
      alert('登录失败: ' + result.msg);
      return;
    }
    // 登录成功时:
    token = result.data.token || '';
    alert('登录成功: ' + result.msg);
  } catch (err) {
    console.error(err);
  }
}

async function handleGetCart() {
  try {
    const resp = await fetch('http://127.0.0.1:8080/cart/get_cart?user_id=123', {
      headers: { 'Authorization': token },
    });
    const result = await resp.json();
    if (result.code !== 0) {
      alert('获取购物车失败: ' + result.msg);
      return;
    }
    document.getElementById('cart').textContent = JSON.stringify(result.data, null, 2);
  } catch (err) {
    console.error(err);
  }
}