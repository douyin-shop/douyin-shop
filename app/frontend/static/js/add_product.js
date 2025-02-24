document.getElementById('productForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const form = event.target;

    let cnt = 0;
    // 获取选中的分类
    const categories = Array.from(document.querySelectorAll('input[name="categories[]"]:checked'))
        .map(c => {
            return {
                id: cnt ++,
                name: c.value
            };
        });

    const productData = {
        name: form.name.value,
        description: form.description.value,
        price: parseFloat(form.price.value),
        category: JSON.stringify(categories)
    };

    const formData = new FormData();
    for (const key in productData) {
        formData.append(key, productData[key]);
    }

    const imageFile = form.image_file.files[0];
    if (!imageFile) {
        alert('请选择商品图片');
        return;
    }
    formData.append('image_file', imageFile);

    fetch('/product/add', {
        method: 'POST',
        body: formData,
        headers: {
            'Authorization': `${localStorage.getItem('authToken')}`
        }
    })
        .then(response => response.json())
        .then(data => {
            document.getElementById('response').innerHTML = `
                <div style="color: #00ff88; padding: 10px;">
                    ✅ 商品添加成功！

                    ${JSON.stringify(data, null, 2)}
                </div>`;
        })
        .catch(error => {
            document.getElementById('response').innerHTML = `
                <div style="color: #ff0050; padding: 10px;">
                    ❌ 错误：${error}
                </div>`;
        });
});

// 显示文件名
document.getElementById('image_file').addEventListener('change', function(e) {
    document.getElementById('file-name').textContent = e.target.files[0].name;
});