import { useState, useEffect } from 'react';
import { api } from './api/api';

function App() {
    const [userId, setUserId] = useState(null);
    const [user, setUser] = useState(null);
    const [transactions, setTransactions] = useState([]);
    const [stats, setStats] = useState({ total_income: 0, total_expense: 0, balance: 0 });
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    // Форма логина
    const [loginId, setLoginId] = useState('');

    // Форма транзакции
    const [form, setForm] = useState({ type: 'income', amount: '', category: '', description: '' });

    // Форма редактирования
    const [editId, setEditId] = useState(null);
    const [editForm, setEditForm] = useState({ type: 'income', amount: '', category: '', description: '' });

    // Админка
    const [adminUsers, setAdminUsers] = useState([]);
    const [showAdmin, setShowAdmin] = useState(false);

    const loadData = async (uid) => {
        setLoading(true);
        setError('');
        try {
            const [userData, txns, statData] = await Promise.all([
                api.getUser(uid),
                api.getTransactions(uid),
                api.getStatistics(uid),
            ]);
            setUser(userData);
            setTransactions(txns || []);
            setStats(statData || { total_income: 0, total_expense: 0, balance: 0 });
            setUserId(uid);
        } catch (err) {
            setError(err.message);
            setUser(null);
            setTransactions([]);
            setStats({ total_income: 0, total_expense: 0, balance: 0 });
        } finally {
            setLoading(false);
        }
    };

    const handleLogin = async (e) => {
        e.preventDefault();
        const id = parseInt(loginId);
        if (isNaN(id)) {
            setError('Введите корректный ID');
            return;
        }
        await loadData(id);
    };

    const handleLogout = () => {
        setUserId(null);
        setUser(null);
        setTransactions([]);
        setStats({ total_income: 0, total_expense: 0, balance: 0 });
        setLoginId('');
        setError('');
        setShowAdmin(false);
    };

    const handleCreateTransaction = async (e) => {
        e.preventDefault();
        if (!form.amount || !form.category) {
            setError('Сумма и категория обязательны');
            return;
        }
        try {
            await api.createTransaction(userId, { ...form, amount: parseFloat(form.amount) });
            setForm({ type: 'income', amount: '', category: '', description: '' });
            await loadData(userId);
        } catch (err) {
            setError(err.message);
        }
    };

    const handleEdit = (txn) => {
        setEditId(txn.id);
        setEditForm({ type: txn.type, amount: txn.amount.toString(), category: txn.category, description: txn.description });
    };

    const handleUpdate = async (e) => {
        e.preventDefault();
        try {
            await api.updateTransaction(userId, editId, { ...editForm, amount: parseFloat(editForm.amount) });
            setEditId(null);
            await loadData(userId);
        } catch (err) {
            setError(err.message);
        }
    };

    const handleDelete = async (id) => {
        if (!confirm('Удалить транзакцию?')) return;
        try {
            await api.deleteTransaction(userId, id);
            await loadData(userId);
        } catch (err) {
            setError(err.message);
        }
    };

    const handleDeleteUser = async (id) => {
        if (!confirm('Удалить пользователя и все его транзакции?')) return;
        try {
            await api.deleteUser(id);
            setAdminUsers(adminUsers.filter(u => u.id !== id));
        } catch (err) {
            setError(err.message);
        }
    };


    const [regUsername, setRegUsername] = useState('');
    const [regPassword, setRegPassword] = useState('');
    const [showRegister, setShowRegister] = useState(false);

    const handleRegister = async (e) => {
        e.preventDefault();
        if (!regUsername || !regPassword) {
            setError('Заполните все поля');
            return;
        }
        try {
            const newUser = await api.createUser(regUsername, regPassword);
            setRegUsername('');
            setRegPassword('');
            setShowRegister(false);
            setError('');
            alert(`Пользователь создан! Ваш ID: ${newUser.id}`);
        } catch (err) {
            setError(err.message);
        }
    };

    if (!userId) {
        return (
            <div style={{ maxWidth: 400, margin: '100px auto', textAlign: 'center' }}>
                <h1>💰 Finance App</h1>

                {!showRegister ? (
                    <>
                        <form onSubmit={handleLogin}>
                            <input type="number" placeholder="Введите ID пользователя" value={loginId}
                                   onChange={(e) => setLoginId(e.target.value)}
                                   style={{ padding: 10, fontSize: 16, width: '100%', marginBottom: 10 }} />
                            <button type="submit" style={{ padding: 10, fontSize: 16, width: '100%', cursor: 'pointer' }}>
                                Войти
                            </button>
                        </form>
                        <p style={{ marginTop: 15 }}>
                            Нет аккаунта?{' '}
                            <button onClick={() => setShowRegister(true)} style={{ background: 'none', border: 'none', color: '#3498db', cursor: 'pointer', textDecoration: 'underline' }}>
                                Зарегистрироваться
                            </button>
                        </p>
                    </>
                ) : (
                    <>
                        <form onSubmit={handleRegister}>
                            <input type="text" placeholder="Имя пользователя" value={regUsername}
                                   onChange={(e) => setRegUsername(e.target.value)}
                                   style={{ padding: 10, fontSize: 16, width: '100%', marginBottom: 10 }} />
                            <input type="password" placeholder="Пароль" value={regPassword}
                                   onChange={(e) => setRegPassword(e.target.value)}
                                   style={{ padding: 10, fontSize: 16, width: '100%', marginBottom: 10 }} />
                            <button type="submit" style={{ padding: 10, fontSize: 16, width: '100%', cursor: 'pointer' }}>
                                Зарегистрироваться
                            </button>
                        </form>
                        <p style={{ marginTop: 15 }}>
                            Уже есть аккаунт?{' '}
                            <button onClick={() => setShowRegister(false)} style={{ background: 'none', border: 'none', color: '#3498db', cursor: 'pointer', textDecoration: 'underline' }}>
                                Войти
                            </button>
                        </p>
                    </>
                )}
                {error && <p style={{ color: 'red' }}>{error}</p>}
            </div>
        );
    }

    return (
        <div style={{ maxWidth: 800, margin: '20px auto', padding: '0 20px' }}>
            <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <h1>💰 Finance App</h1>
                <div>
                    <button onClick={() => setShowAdmin(!showAdmin)} style={{ marginRight: 10 }}>
                        {showAdmin ? 'Назад' : 'Админка'}
                    </button>
                    <button onClick={handleLogout}>Выйти</button>
                </div>
            </div>

            {error && <p style={{ color: 'red', padding: 10, background: '#fee' }}>{error}</p>}

            {showAdmin ? (
                <div>
                    <h2>Управление пользователями</h2>
                    {adminUsers.length === 0 && <p>Загрузите пользователей или попробуйте удалить по ID</p>}
                    {adminUsers.map(u => (
                        <div key={u.id} style={{ display: 'flex', justifyContent: 'space-between', padding: 10, border: '1px solid #ddd', marginBottom: 5 }}>
                            <span>{u.username} (ID: {u.id})</span>
                            <button onClick={() => handleDeleteUser(u.id)} style={{ background: '#e74c3c', color: 'white', border: 'none', padding: '5px 10px', cursor: 'pointer' }}>
                                Удалить
                            </button>
                        </div>
                    ))}
                </div>
            ) : (
                <>
                    <div style={{ display: 'flex', gap: 20, marginBottom: 20 }}>
                        <p><strong>Пользователь:</strong> {user?.username} (ID: {userId})</p>
                        <p>🟢 Доходы: <strong>{stats.total_income} ₽</strong></p>
                        <p>🔴 Расходы: <strong>{stats.total_expense} ₽</strong></p>
                        <p>💰 Баланс: <strong style={{ color: stats.balance >= 0 ? 'green' : 'red' }}>{stats.balance} ₽</strong></p>
                    </div>

                    <div style={{ border: '1px solid #ddd', padding: 15, marginBottom: 20 }}>
                        <h3>Добавить транзакцию</h3>
                        <form onSubmit={handleCreateTransaction} style={{ display: 'grid', gridTemplateColumns: '1fr 1fr 1fr 1fr', gap: 10 }}>
                            <select value={form.type} onChange={(e) => setForm({ ...form, type: e.target.value })}>
                                <option value="income">Доход</option>
                                <option value="expense">Расход</option>
                            </select>
                            <input type="number" placeholder="Сумма" value={form.amount} onChange={(e) => setForm({ ...form, amount: e.target.value })} />
                            <input type="text" placeholder="Категория" value={form.category} onChange={(e) => setForm({ ...form, category: e.target.value })} />
                            <button type="submit">Добавить</button>
                            <input type="text" placeholder="Описание (необязательно)" value={form.description} onChange={(e) => setForm({ ...form, description: e.target.value })} style={{ gridColumn: 'span 3' }} />
                        </form>
                    </div>

                    <h3>Транзакции</h3>
                    {loading ? (
                        <p>Загрузка...</p>
                    ) : transactions.length === 0 ? (
                        <p>Нет транзакций</p>
                    ) : (
                        transactions.map((txn) => (
                            <div key={txn.id} style={{ border: `2px solid ${txn.type === 'income' ? '#2ecc71' : '#e74c3c'}`, padding: 10, marginBottom: 8 }}>
                                {editId === txn.id ? (
                                    <form onSubmit={handleUpdate} style={{ display: 'grid', gridTemplateColumns: '1fr 1fr 1fr 1fr', gap: 10 }}>
                                        <select value={editForm.type} onChange={(e) => setEditForm({ ...editForm, type: e.target.value })}>
                                            <option value="income">Доход</option>
                                            <option value="expense">Расход</option>
                                        </select>
                                        <input type="number" value={editForm.amount} onChange={(e) => setEditForm({ ...editForm, amount: e.target.value })} />
                                        <input type="text" value={editForm.category} onChange={(e) => setEditForm({ ...editForm, category: e.target.value })} />
                                        <div>
                                            <button type="submit" style={{ marginRight: 5 }}>✓</button>
                                            <button type="button" onClick={() => setEditId(null)}>✕</button>
                                        </div>
                                    </form>
                                ) : (
                                    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                                        <div>
                                            <strong style={{ color: txn.type === 'income' ? '#27ae60' : '#c0392b' }}>
                                                {txn.type === 'income' ? '+' : '-'}{txn.amount} ₽
                                            </strong>
                                            {' | '}{txn.category}
                                            {txn.description && ` | ${txn.description}`}
                                            <br />
                                            <small>{new Date(txn.created_at).toLocaleString()}</small>
                                        </div>
                                        <div>
                                            <button onClick={() => handleEdit(txn)} style={{ marginRight: 5 }}>✎</button>
                                            <button onClick={() => handleDelete(txn.id)} style={{ background: '#e74c3c', color: 'white', border: 'none' }}>✕</button>
                                        </div>
                                    </div>
                                )}
                            </div>
                        ))
                    )}
                </>
            )}
        </div>
    );
}

export default App;