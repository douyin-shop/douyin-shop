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

    // 构造查询参数（根据实际需要补充签名相关参数）
    const queryParams = new URLSearchParams({
        order_id: orderId,                // 订单ID
        transaction_id: transactionId,    // 交易ID
        payment_time: new Date().toISOString(), // 自动生成支付时间
        sign: 'mock_signature_placeholder',     // 签名占位符（需根据实际签名逻辑生成）
        sign_type: 'HMAC-SHA256'          // 签名类型
    });

    fetch(`/payment/callback?${queryParams}`, {
        method: 'GET',
        headers: {
            'Authorization': `${localStorage.getItem('authToken')}`
        }
    })
        .then(response => {
            if (response.status === 401) {
                window.location.href = 'index.html';
                return;
            }
            return response.json();
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