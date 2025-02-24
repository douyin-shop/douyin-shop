document.addEventListener('DOMContentLoaded', () => {
    if (!localStorage.getItem('authToken')) {
        window.location.href = 'index.html';
        return;
    }

    const urlParams = new URLSearchParams(window.location.search);
    const transactionId = urlParams.get('transaction_id') || 'Unknown';
    const orderId = urlParams.get('order_id') || 'Unknown';

    document.getElementById('transactionId').textContent = transactionId;
    document.getElementById('orderId').textContent = orderId;
});

function simulatePaymentSuccess() {
    const transactionId = document.getElementById('transactionId').textContent;
    const orderId = document.getElementById('orderId').textContent;

    fetch('/payment/callback', {
        method: 'GET',  // Changed to GET based on the router configuration
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `${localStorage.getItem('authToken')}`
        }
    })
        .then(response => {
            // 如果状态码是401，说明用户未登录，跳转到登录页面
            if (response.status === 401) {
                window.location.href = 'index.html';
                return;
            }

            return response.json()
        })
        .then(result => {
            document.getElementById('paymentMessage').style.display = 'block';
            if (result.code === 0 && result.data.success) {
                setTimeout(() => {
                    window.location.href = 'order_placeholder.html';
                }, 2000);
            } else {
                throw new Error(result.msg || 'Payment failed');
            }
        })
        .catch(error => {
            console.error('Error:', error);
            alert('Payment failed: ' + error.message);
            document.getElementById('paymentMessage').style.display = 'none';
        });
}