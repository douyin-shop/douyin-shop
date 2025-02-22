// 保持原有的JavaScript逻辑不变
const API_BASE = '';

document.addEventListener('DOMContentLoaded', async () => {
    if (!localStorage.getItem('authToken')) {
        window.location.href = 'index.html';
    }
    await loadProducts();
});

async function loadProducts() {
    try {
        const response = await fetch(`${API_BASE}/product/list`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `${localStorage.getItem('authToken')}`
            },
            body: JSON.stringify({ page: 1, pageSize: 20 })
        });

        const data = await response.json();
        renderProducts(data.data.products);
    } catch (error) {
        console.error('Error loading products:', error);
    }
}

function renderProducts(products) {
    const container = document.getElementById('productList');
    container.innerHTML = products.map(product => `
            <div class="product-card">
                <img src="${product.picture}" class="product-image" alt="${product.name}">
                <div class="product-info">
                    <div class="product-title">${product.name}</div>
                    <div class="product-price">¥${product.price.toFixed(2)}</div>
                </div>
                <a href="javascript:void(0)" onclick="viewProductDetail(${product.id})" class="view-button">
                    查看详情
                </a>
            </div>
        `).join('');
}

function viewProductDetail(productId) {
    window.location.href = `product-detail.html?id=${productId}`;
}

async function searchProducts() {
    const query = document.getElementById('searchInput').value;
    try {
        const response = await fetch(`${API_BASE}/product/search`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `${localStorage.getItem('authToken')}`
            },
            body: JSON.stringify({ query })
        });

        const data = await response.json();
        if(data.code === -1 || data.data.results.length === 0){
            alert('未找到相关商品');
            return;
        }
        renderProducts(data.data.results);
    } catch (error) {
        console.error('Search error:', error);
    }
}

// 添加搜索输入框的实时监听
document.getElementById('searchInput').addEventListener('keypress', (e) => {
    if (e.key === 'Enter') {
        searchProducts();
    }
});