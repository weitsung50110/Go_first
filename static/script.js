// JavaScript 檔案
document.getElementById("textForm").addEventListener("submit", function(event) {
    event.preventDefault(); // 阻止表單預設提交行為

    const formData = {
        // 獲取表單中的數據並添加到 FormData 對象中
        inputText: document.getElementById("inputText").value, // 獲取表單輸入框中的文字,
        inputName: document.getElementById('inputName').value
    }
    console.log("123",formData)

    const requestOptions = {
        method: "POST", // 使用 POST 方法
        body: new URLSearchParams(formData), // 提交表單文字
        headers: {
            "Content-Type": "application/x-www-form-urlencoded" // 設置標頭內容類型
        }
    }

    fetch("/submit",requestOptions)
    .then(response => response.json()) // 獲取伺服器返回的文字內容
    .then(data => {
        document.getElementById("textDisplay").innerHTML =data.inputName+" : "+data.inputText; 
        // 將文字內容顯示在 id 為 textDisplay 的區域
    });
});

