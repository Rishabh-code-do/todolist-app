import React, { useState, useEffect } from 'react';
import './App.css';

const API_BASE = process.env.REACT_APP_API_BASE;

function App() {
  const [todos, setTodos] = useState([]);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [dueDate, setDueDate] = useState('');
  const [editingId, setEditingId] = useState(null);

  const fetchTodos = async () => {
    try {
      const res = await fetch(`${API_BASE}/get`);
      const data = await res.json();
      setTodos(data);
    } catch (err) {
      console.error('Failed to fetch todos:', err);
    }
  };

  const resetForm = () => {
    setTitle('');
    setDescription('');
    setDueDate('');
    setEditingId(null);
  };

  const addOrUpdateTodo = async () => {
    if (!title.trim()) return;

    const body = JSON.stringify({
      title,
      description,
      duedate: new Date(dueDate).toISOString(),
    });

    const url = editingId ? `${API_BASE}/update/${editingId}` : `${API_BASE}/create`;
    const method = editingId ? 'PATCH' : 'POST';

    try {
      await fetch(url, {
        method,
        headers: { 'Content-Type': 'application/json' },
        body,
      });
      resetForm();
      fetchTodos();
    } catch (err) {
      console.error('Failed to add or update todo:', err);
    }
  };

  const editTodo = async (id) => {
    try {
      const res = await fetch(`${API_BASE}/get/${id}`);
      const data = await res.json();

      setTitle(data.title);
      setDescription(data.description);
      const formattedDate = data.duedate
      ? new Date(data.duedate).toISOString().split('T')[0]
      : '';
      setDueDate(formattedDate);
      setEditingId(id);
      }catch (err) {
      console.error('Failed to fetch todo:', err);
    }
  };

  const completeTodo = async (id) => {
    try {
      await fetch(`${API_BASE}/complete/${id}`, { method: 'PATCH' });
      fetchTodos();
    } catch (err) {
      console.error('Failed to complete todo:', err);
    }
  };

  const deleteTodo = async (id) => {
    try {
      await fetch(`${API_BASE}/delete/${id}`, { method: 'DELETE' });
      fetchTodos();
    } catch (err) {
      console.error('Failed to delete todo:', err);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  return (
    <div className="App">
      <h1>TODO List</h1>

      <input
        type="text"
        value={title}
        placeholder="Title"
        onChange={(e) => setTitle(e.target.value)}
      /><br />

      <textarea
        value={description}
        placeholder="Description"
        onChange={(e) => setDescription(e.target.value)}
      /><br />

      <input
        type="date"
        value={dueDate}
        onChange={(e) => setDueDate(e.target.value)}
      /><br />

      <button onClick={addOrUpdateTodo}>
        {editingId ? 'Update' : 'Add'} Todo
      </button>
      {editingId && <button onClick={resetForm}>Cancel</button>}

      <hr />
      <ul>
        {todos.map((t) => (
          <li key={t.id}>
            <strong>{t.title}</strong>
            <p>{t.description}</p>
            <p>
              Due:{' '}
              {t.duedate
                ? new Intl.DateTimeFormat('en-US', {
                    month: 'short',
                    day: '2-digit',
                    year: 'numeric',
                  }).format(new Date(t.duedate))
                : 'No due date'}
            </p>
            <button onClick={() => editTodo(t.id)}>Edit</button>
            <button onClick={() => completeTodo(t.id)}>Complete</button>
            <button onClick={() => deleteTodo(t.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
