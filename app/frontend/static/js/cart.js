const API_BASE = '';

document.addEventListener('DOMContentLoaded', async () => {
    if (!localStorage.getItem('authToken')) {
        window.location.href = 'index.html';
    }
    await loadCart();
});

async function loadCart() {
    try {
        const response = await fetch(`${API_BASE}/cart/get_all`, {
            method: 'GET',
            headers: {
                'Authorization': `${localStorage.getItem('authToken')}`
            }
        });


        // 如果状态码是401，说明用户未登录，跳转到登录页面
        if (response.status === 401) {
            window.location.href = 'index.html';
            return;
        }


        const data = await response.json();
        renderCart(data.data.cart);
    } catch (error) {
        console.error('Error loading cart:', error);
    }
}

function renderCart(cart) {
    const container = document.getElementById('cartItems');
    container.innerHTML = cart.items.map(item => `
            <div class="cart-item">
                <img src="${item.product.imageUrl}" class="item-image" alt="${item.product.name}">
                <div class="item-info">
                    <div class="item-name">${item.product.name}</div>
                    <div class="item-price">¥${item.product.price.toFixed(2)}</div>
                    <div class="quantity-control">
                        <button class="quantity-btn" onclick="updateQuantity(${item.product.id},-1)">-</button>
                        <input type="number" class="quantity-input" value="${item.quantity}" 
                               min="1" onchange="updateQuantity(${item.product.id}, this.value - item.quantity)">
                        <button class="quantity-btn" onclick="updateQuantity(${item.product.id}, 1)">+</button>
                    </div>
                </div>
            </div>
        `).join('');

    document.getElementById('totalPrice').textContent = cart.total_price.toFixed(2);
}

async function updateQuantity(productId, editNumber) {
    try {
        const response = await fetch(`${API_BASE}/cart/add`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': localStorage.getItem('authToken')
            },
            body: JSON.stringify({
                product_id: productId,
                product_num: parseInt(editNumber)
            })
        });

        // 如果状态码是401，说明用户未登录，跳转到登录页面
        if (response.status === 401) {
            window.location.href = 'index.html';
            return;
        }

        // 如果返回结果中的code为-1，说明商品数量不足
        const result = await response.json();
        if (result.code === -1) {
            alert(result.msg);
            return;
        }

        if (response.ok) {
            await loadCart();
        }
    } catch (error) {
        console.error('Error updating quantity:', error);
    }
}

async function proceedToCheckout() {
    try {
        const response = await fetch(`${API_BASE}/cart/get_all`, {
            headers: {
                'Authorization': localStorage.getItem('authToken')
            }
        });

        // 如果状态码是401，说明用户未登录，跳转到登录页面
        if (response.status === 401) {
            window.location.href = 'index.html';
            return;
        }


        const cart = await response.json();
        if (cart.data.cart.items.length === 0) {
            alert('购物车为空！');
            return;
        }
        window.location.href = 'checkout.html';
    } catch (error) {
        console.error('Error:', error);
    }
}