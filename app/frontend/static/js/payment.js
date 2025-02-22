document.addEventListener('DOMContentLoaded', () => {
    if (!localStorage.getItem('authToken')) {
        window.location.href = 'index.html';
    }

    const urlParams = new URLSearchParams(window.location.search);
    document.getElementById('transactionId').textContent = urlParams.get('transaction_id');
    document.getElementById('orderId').textContent = urlParams.get('order_id');
});

function simulatePaymentSuccess() {
    document.getElementById('paymentMessage').style.display = 'block';
    setTimeout(() => {
        // 这里暂时跳转到占位页面，等订单页面完成后替换成实际URL
        window.location.href = 'order_placeholder.html';
    }, 2000);
}