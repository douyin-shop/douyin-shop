const API_BASE = '';

document.addEventListener('DOMContentLoaded', async () => {
    const productId = new URLSearchParams(window.location.search).get('id');
    if (productId) {
        await loadProduct(productId);
    }
});

async function loadProduct(id) {
    try {
        const response = await fetch(`${API_BASE}/product/get`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `${localStorage.getItem('authToken')}`
            },
            body: JSON.stringify({ id: parseInt(id) })
        });

        const data = await response.json();
        renderProduct(data.data.product);
    } catch (error) {
        console.error('Error loading product:', error);
    }
}

function renderProduct(product) {
    document.getElementById('productName').textContent = product.name;
    document.getElementById('productImage').src = product.picture;
    document.getElementById('productDescription').textContent = product.description;
    document.getElementById('productPrice').textContent = `¥${product.price.toFixed(2)}`;
}

async function addToCart() {
    const quantity = parseInt(document.getElementById('quantity').value) || 1;
    try {
        const response = await fetch(`${API_BASE}/cart/add`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `${localStorage.getItem('authToken')}`
            },
            body: JSON.stringify({
                product_id: parseInt(new URLSearchParams(window.location.search).get('id')),
                product_num: quantity
            })
        });

        if (response.ok) {
            alert('成功加入购物车 🛒');
        } else {
            const error = await response.json();
            alert('操作失败: ' + error.message);
        }
    } catch (error) {
        console.error('添加购物车失败:', error);
    }
}