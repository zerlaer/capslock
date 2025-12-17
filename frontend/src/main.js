import './style.css';
import './app.css';

import { StartPreventLock, StopPreventLock, GetStatus, SetInterval } from '../wailsjs/go/main/App';

// DOM元素
const statusElement = document.getElementById('status');
const statusInfoElement = document.getElementById('statusInfo');
const intervalInput = document.getElementById('interval');
const startBtn = document.getElementById('startBtn');
const stopBtn = document.getElementById('stopBtn');
const historyElement = document.getElementById('history');

// 初始化
window.addEventListener('DOMContentLoaded', () => {
    updateStatus();
    addToHistory('应用已启动');
});

// 获取当前状态并更新UI
async function updateStatus() {
    try {
        const status = await GetStatus();
        updateUI(status);
    } catch (err) {
        console.error('获取状态失败:', err);
        addToHistory('获取状态失败: ' + err.message);
    }
}

// 更新UI
function updateUI(status) {
    const { running, interval } = status;
    
    // 更新状态显示
    statusElement.textContent = running ? '运行中' : '已停止';
    statusElement.className = 'status ' + (running ? 'running' : 'stopped');
    
    
    // 更新按钮状态
    startBtn.disabled = running;
    stopBtn.disabled = !running;
    
    // 更新时间间隔输入框
    intervalInput.value = interval;
}

// 启动防止锁屏
window.startPreventLock = async function() {
    try {
        // 获取时间间隔
        const interval = parseInt(intervalInput.value);
        
        // 设置时间间隔
        await SetInterval(interval);
        
        // 启动防止锁屏
        const result = await StartPreventLock();
        
        // 更新状态和历史记录
        await updateStatus();
        addToHistory(result);
    } catch (err) {
        console.error('启动失败:', err);
        addToHistory('启动失败: ' + err.message);
    }
};

// 停止防止锁屏
window.stopPreventLock = async function() {
    try {
        const result = await StopPreventLock();
        
        // 更新状态和历史记录
        await updateStatus();
        addToHistory(result);
    } catch (err) {
        console.error('停止失败:', err);
        addToHistory('停止失败: ' + err.message);
    }
};

// 添加到历史记录
function addToHistory(message) {
    const now = new Date();
    const timeString = now.toLocaleTimeString();
    
    const historyItem = document.createElement('div');
    historyItem.className = 'history-item';
    historyItem.innerHTML = `<span class="time">[${timeString}]</span> ${message}`;
    
    historyElement.appendChild(historyItem);
    
    // 滚动到底部
    historyElement.scrollTop = historyElement.scrollHeight;
    
    // 限制历史记录数量
    if (historyElement.children.length > 20) {
        historyElement.removeChild(historyElement.firstChild);
    }
};

// 监听时间间隔输入框变化
intervalInput.addEventListener('change', async () => {
    try {
        const interval = parseInt(intervalInput.value);
        if (interval <= 0) {
            intervalInput.value = 30; // 重置为默认值
            return;
        }
        
        // 设置时间间隔
        const result = await SetInterval(interval);
        
        // 更新状态和历史记录
        await updateStatus();
        addToHistory(result);
    } catch (err) {
        console.error('设置时间间隔失败:', err);
        addToHistory('设置时间间隔失败: ' + err.message);
    }
});
