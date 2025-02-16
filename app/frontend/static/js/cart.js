// const API_BASE = 'http://127.0.0.1:8080';
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
    const tbody = document.getElementById('cartItems');
    tbody.innerHTML = cart.items.map(item => `
        <tr>
            <td>${item.product.name}</td>
            <td>$${item.product.price.toFixed(2)}</td>
            <td>${item.quantity}</td>
            <td>$${(item.product.price * item.quantity).toFixed(2)}</td>
        </tr>
    `).join('');

    document.getElementById('totalPrice').textContent = cart.total_price.toFixed(2);
}