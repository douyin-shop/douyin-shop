const API_BASE = ''; // http://localhost:8080';
document.addEventListener('DOMContentLoaded', async () => {
    await loadOrders();
});

async function loadOrders() {
    try {
        const response = await fetch(`${API_BASE}/order/list`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ page: 1, pageSize: 20 })
        });
        const data = await response.json();
        if(data.code === -1){
            alert(data.msg);
            return;
        }
        renderOrders(data.data.orders);
    } catch (error) {
        console.error('Error loading orders:', error);
    }
}

function renderOrders(orders) {
    const container = document.getElementById('orderGrid');
    container.innerHTML = orders.map(order => `
        <div class="order-card">
            <h3>Order ID: ${order.id}</h3>
            <p>User ID: ${order.user_id}</p>
            <p>Total Price: $${order.total_price}</p>
            <p>Created At: ${new Date(order.created_at * 1000).toLocaleString()}</p>
            <button onclick="viewOrderDetail(${order.id})">View Details</button>
        </div>
    `).join('');
}

function viewOrderDetail(orderId) {
    window.location.href = `order-detail.html?id=${orderId}`; // Assuming you have an order detail page
}
