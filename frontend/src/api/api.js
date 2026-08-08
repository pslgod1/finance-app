const API_URL = 'http://localhost:8080/api';

export const api = {
    // Пользователи
    createUser: async (username, password) => {
        const res = await fetch(`${API_URL}/users`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, password }),
        });
        if (!res.ok) throw new Error((await res.json()).message);
        return res.json();
    },

    getUser: async (id) => {
        const res = await fetch(`${API_URL}/users/${id}`);
        if (!res.ok) throw new Error('Пользователь не найден');
        return res.json();
    },

    // Транзакции
    getTransactions: async (userId) => {
        const res = await fetch(`${API_URL}/users/${userId}/transactions`);
        if (!res.ok) throw new Error('Ошибка загрузки');
        return res.json();
    },

    createTransaction: async (userId, data) => {
        const res = await fetch(`${API_URL}/users/${userId}/transactions`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data),
        });
        if (!res.ok) throw new Error((await res.json()).message);
        return res.json();
    },

    updateTransaction: async (userId, id, data) => {
        const res = await fetch(`${API_URL}/users/${userId}/transactions/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data),
        });
        if (!res.ok) throw new Error((await res.json()).message);
        return res.json();
    },

    deleteTransaction: async (userId, id) => {
        const res = await fetch(`${API_URL}/users/${userId}/transactions/${id}`, {
            method: 'DELETE',
        });
        if (!res.ok) throw new Error('Ошибка удаления');
    },

    // Статистика
    getStatistics: async (userId) => {
        const res = await fetch(`${API_URL}/users/${userId}/statistics`);
        if (!res.ok) throw new Error('Ошибка загрузки');
        return res.json();
    },

    // Админка
    deleteUser: async (id) => {
        const res = await fetch(`${API_URL}/admin/users/${id}`, {
            method: 'DELETE',
        });
        if (!res.ok) throw new Error((await res.json()).message);
    },
};