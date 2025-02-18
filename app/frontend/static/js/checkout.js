// document.addEventListener('DOMContentLoaded', () => {
//     if (!localStorage.getItem('authToken')) {
//         window.location.href = 'index.html';
//     }
//
//     // 填充年份选项
//     const yearSelect = document.getElementById('credit_card_expiration_year');
//     const currentYear = new Date().getFullYear();
//     for (let i = 0; i < 5; i++) {
//         const option = document.createElement('option');
//         option.value = currentYear + i;
//         option.textContent = currentYear + i;
//         yearSelect.appendChild(option);
//     }
//
//     document.getElementById('checkoutForm').addEventListener('submit', async (e) => {
//         e.preventDefault();
//
//         const formData = {
//             firstname: document.getElementById('firstname').value,
//             lastname: document.getElementById('lastname').value,
//             email: document.getElementById('email').value,
//             address: {
//                 street_address: document.getElementById('street_address').value,
//                 city: document.getElementById('city').value,
//                 state: document.getElementById('state').value,
//                 country: document.getElementById('country').value,
//                 zip_code: document.getElementById('zip_code').value
//             },
//             credit_card: {
//                 credit_card_number: document.getElementById('credit_card_number').value,
//                 credit_card_cvv: parseInt(document.getElementById('credit_card_cvv').value),
//                 credit_card_expiration_month: parseInt(document.getElementById('credit_card_expiration_month').value),
//                 credit_card_expiration_year: parseInt(document.getElementById('credit_card_expiration_year').value)
//             }
//         };
//
//         try {
//             const response = await fetch('/checkout', {
//                 method: 'POST',
//                 headers: {
//                     'Content-Type': 'application/json',
//                     'Authorization': localStorage.getItem('authToken')
//                 },
//                 body: JSON.stringify(formData)
//             });
//
//             const data = await response.json();
//
//
//             if (data.code===0) {
//                 window.location.href = `payment.html?transaction_id=${data.data.transaction_id}&order_id=${data.data.order_id}`;
//             } else {
//                 alert(`Checkout failed: ${data.data.message}`);
//             }
//         } catch (error) {
//             console.error('Checkout error:', error);
//             alert('Checkout process failed');
//         }
//     });
// });

document.addEventListener('DOMContentLoaded', () => {
    if (!localStorage.getItem('authToken')) {
        window.location.href = 'index.html';
    }

    // 填充月份
    const monthSelect = document.getElementById('credit_card_expiration_month');
    for (let i = 1; i <= 12; i++) {
        const option = document.createElement('option');
        option.value = i;
        option.textContent = i.toString().padStart(2, '0');
        if(i === (new Date().getMonth() + 1)) option.selected = true;
        monthSelect.appendChild(option);
    }

    // 填充年份（未来5年）
    const yearSelect = document.getElementById('credit_card_expiration_year');
    const currentYear = new Date().getFullYear();
    for (let i = 0; i < 5; i++) {
        const option = document.createElement('option');
        const year = currentYear + i;
        option.value = year;
        option.textContent = year;
        if(i === 1) option.selected = true;
        yearSelect.appendChild(option);
    }

    document.getElementById('checkoutForm').addEventListener('submit', async (e) => {
        e.preventDefault();

        const formData = {
            firstname: document.getElementById('firstname').value,
            lastname: document.getElementById('lastname').value,
            email: document.getElementById('email').value,
            address: {
                street_address: document.getElementById('street_address').value,
                city: document.getElementById('city').value,
                state: document.getElementById('state').value,
                country: document.getElementById('country').value,
                zip_code: document.getElementById('zip_code').value
            },
            credit_card: {
                credit_card_number: document.getElementById('credit_card_number').value,
                credit_card_cvv: parseInt(document.getElementById('credit_card_cvv').value),
                credit_card_expiration_month: parseInt(document.getElementById('credit_card_expiration_month').value),
                credit_card_expiration_year: parseInt(document.getElementById('credit_card_expiration_year').value)
            }
        };
        try {
            const response = await fetch('/checkout', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': localStorage.getItem('authToken')
                },
                body: JSON.stringify(formData)
            });

            const data = await response.json();

            if (data.code===0) {
                window.location.href = `payment.html?transaction_id=${data.data.transaction_id}&order_id=${data.data.order_id}`;
            } else {
                alert(`Checkout failed: ${data.data.message}`);
            }
        } catch (error) {
            console.error('结算错误:', error);
            alert('结算失败，请稍后重试');
        }
    });
});