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
                <img src="${item.product.picture}" class="item-image" alt="${item.product.name}">
                <div class="item-info">
                    <div class="item-name">${item.product.name}</div>
                    <div class="item-price">¥${item.product.price.toFixed(2)}</div>
                    <div class="quantity-control">
                        <button class="quantity-btn" onclick="updateQuantity(${item.product.id}, ${item.quantity - 1})">-</button>
                        <input type="number" class="quantity-input" value="${item.quantity}" 
                               min="1" onchange="updateQuantity(${item.product.id}, this.value)">
                        <button class="quantity-btn" onclick="updateQuantity(${item.product.id}, ${item.quantity + 1})">+</button>
                    </div>
                </div>
            </div>
        `).join('');

    document.getElementById('totalPrice').textContent = cart.total_price.toFixed(2);
}

async function updateQuantity(productId, newQuantity) {
    try {
        const response = await fetch(`${API_BASE}/cart/update`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': localStorage.getItem('authToken')
            },
            body: JSON.stringify({
                product_id: productId,
                quantity: parseInt(newQuantity)
            })
        });

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