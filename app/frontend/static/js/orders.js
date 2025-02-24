const API_BASE = ''; // http://localhost:8080';
document.addEventListener('DOMContentLoaded', async () => {
    await loadOrders();
});

async function loadOrders() {
    try {
        const response = await fetch(`${API_BASE}/order/list`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `${localStorage.getItem('authToken')}`
            },
            body: JSON.stringify({ page: 1, pageSize: 20 })
        });
        // 如果状态码是401，说明用户未登录，跳转到登录页面
        if (response.status === 401) {
            window.location.href = 'index.html';
            return;
        }

        const data = await response.json();
        if(data.code === -1){
            alert(data.msg);
            return;
        }
        console.log(data);
        renderOrders(data.data.orders);
    } catch (error) {
        console.error('Error loading orders:', error);
    }
}

// 修改后的渲染函数
function renderOrders(orders) {
    const container = document.getElementById('orderGrid');
    container.innerHTML = orders.map(order => {
        const totalCost = order.order_items.reduce((sum, item) => sum + item.cost, 0);
        return `
                <div class="order-card">
                    <div class="order-header">
                        <div>
                            <span class="order-status">待发货</span>
                            <span class="order-time">${new Date(order.created_at * 1000).toLocaleString()}</span>
                        </div>
                        <span class="order-id">订单号：${order.order_id}</span>
                    </div>

                    ${order.order_items.map(item => `
                        <div class="product-item">
                            <div class="product-image"></div>
                            <div class="product-info">
                                <div class="product-title">商品名称（需要接口补充）</div>
                                <div class="product-spec">
                                    <span>规格：默认</span>
                                    <span style="margin: 0 8px">|</span>
                                    <span>数量：${item.item.quantity}</span>
                                </div>
                                <div class="product-price">¥${item.cost.toFixed(2)}</div>
                            </div>
                        </div>
                    `).join('')}

                    <div class="order-footer">
                        <div class="total-price">实付：¥${totalCost.toFixed(2)}</div>
                        <div class="action-buttons">
                            <button class="btn btn-secondary">查看物流</button>
                            <button class="btn btn-primary" onclick="viewOrderDetail('${order.order_id}')">订单详情</button>
                        </div>
                    </div>
                </div>
            `;
    }).join('');
}


function viewOrderDetail(orderId) {
    window.location.href = `order-detail.html?id=${orderId}`; // Assuming you have an order detail page
}
