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
                <img src="${product.imageUrl}" class="product-image" alt="${product.name}">
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
    const searchInput = document.getElementById('searchInput').value;
    try {
        const response = await fetch(`${API_BASE}/product/search`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `${localStorage.getItem('authToken')}`
            },
            body: JSON.stringify({
                searchQuery: {  // 符合Protobuf结构
                    keyWord: searchInput,    // 关键词
                    pageNum: 1,             // 当前页码（可按需动态获取）
                    pageSize: 10,           // 每页数量（可按需动态获取）
                    CategoryName: "",       // 分类（可按需添加选择器获取）
                    minPrice: 0,            // 最低价（可按需添加输入框获取）
                    maxPrice: 0             // 最高价（可按需添加输入框获取）
                }
            })
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
// 添加搜索输入框的防抖监听
let searchTimeout;
document.getElementById('searchInput').addEventListener('input', () => {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
        searchProducts();
    }, 2000);
});