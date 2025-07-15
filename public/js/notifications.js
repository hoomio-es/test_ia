// Toastify.js configuration
const Toastify = window.Toastify;

// Function to show success message
function showSuccess(message) {
    Toastify({
        text: message,
        duration: 3000,
        gravity: "top",
        position: "right",
        backgroundColor: "#4CAF50",
        className: "toast",
    }).showToast();
}

// Function to show error message
function showError(message) {
    Toastify({
        text: message,
        duration: 3000,
        gravity: "top",
        position: "right",
        backgroundColor: "#f44336",
        className: "toast",
    }).showToast();
}

// Loading spinner
const loadingSpinner = document.createElement('div');
loadingSpinner.className = 'loading-spinner';
loadingSpinner.innerHTML = `
    <div class="spinner">
        <div class="bounce1"></div>
        <div class="bounce2"></div>
        <div class="bounce3"></div>
    </div>
`;

// Function to show loading spinner
function showLoadingSpinner() {
    document.querySelector('.search-bar').appendChild(loadingSpinner);
}

// Function to hide loading spinner
function hideLoadingSpinner() {
    loadingSpinner.remove();
}
