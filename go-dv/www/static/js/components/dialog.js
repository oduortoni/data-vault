/*
 * file: go-dv/www/static/js/components/dialogue.js
 * description: This file is used to render the dialog view.
 * author: toni
 * date: 2025-06-28
 * version: 1.0.0
 * license: MIT
 * copyright: 2025 toni
 * contact: oduortoni@gmail.com
 */

const Dialog = (title, message, onConfirm = null, onCancel = null) => {
    let overlay = document.getElementById("confirm-overlay");

    // Inject if not present
    if (!overlay) {
        overlay = document.createElement("div");
        overlay.id = "confirm-overlay";
        overlay.innerHTML = `
            <div class="modal">
                <h3 id="modal-title"></h3>
                <p id="modal-message"></p>
                <div class="actions">
                    <button id="modal-ok">OK</button>
                    <button id="modal-cancel">Cancel</button>
                </div>
            </div>
        `;
        document.body.appendChild(overlay);
    }

    // Set content
    document.getElementById("modal-title").textContent = title;
    document.getElementById("modal-message").textContent = message;
    overlay.style.display = "flex";

    // Handle OK
    document.getElementById("modal-ok").onclick = () => {
        overlay.style.display = "none";
        if (typeof onConfirm === "function") onConfirm();
    };

    // Handle Cancel
    document.getElementById("modal-cancel").onclick = () => {
        overlay.style.display = "none";
        if (typeof onCancel === "function") onCancel();
    };
};

export default Dialog;