import React, { useState } from 'react';
import { Button, TextField } from '@material-ui/core';
import Alert from '@material-ui/lab/Alert';

import api from '../libs/api';

export default function AddBook() {
  const [title, setTitle] = useState('');
  const [isbn, setIsbn] = useState('');
  const [purchaseFrom, setPurchaseFrom] = useState('');
  const [remark, setRemark] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async () => {
    try {
      await api.post('/books', {
        title,
        isbn,
        purchaseFrom,
        remark,
      });
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <form onClick={handleSubmit} className="form-add">
      <h2>添加新书</h2>
      <div className="form-item">
        <TextField
          id="title"
          label="标题"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
      </div>
      <div className="form-item">
        <TextField
          id="isbn"
          label="ISBN"
          value={isbn}
          onChange={(e) => setIsbn(e.target.value)}
        />
      </div>
      <div className="form-item">
        <TextField
          id="purchaseFrom"
          label="购买于"
          value={purchaseFrom}
          onChange={(e) => setPurchaseFrom(e.target.value)}
        />
      </div>
      <div className="form-item">
        <TextField
          multiline
          minRows={2}
          id="remark"
          label="备忘"
          value={remark}
          onChange={(e) => setRemark(e.target.value)}
        />
      </div>
      {error && <Alert className="form-item" severity="error">{error}</Alert>}
      <div className="form-item">
        <Button variant="contained" color="primary">
          添加
        </Button>
      </div>
    </form>
  );
}
