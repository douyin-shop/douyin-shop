// const API_BASE = 'http://127.0.0.1:8080';
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
            <h3>${product.name}</h3>
            <img src="${product.picture}" alt="${product.name}" style="max-width: 200px;">
            <p>$${product.price.toFixed(2)}</p>
            <button onclick="viewProductDetail(${product.id})">View Details</button>
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
                'Authorization': `Bearer ${localStorage.getItem('authToken')}`
            },
            body: JSON.stringify({ query })
        });

        const data = await response.json();
        renderProducts(data.results);
    } catch (error) {
        console.error('Search error:', error);
    }
}