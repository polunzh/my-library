import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import { Button, CircularProgress, TextField } from '@material-ui/core';
import Alert from '@material-ui/lab/Alert';

import api from '../libs/api';

export default function AddBook() {
  const history = useHistory();
  const [submitting, setSubmitting] = useState(false);
  const [title, setTitle] = useState('');
  const [isbn, setIsbn] = useState('');
  const [purchaseFrom, setPurchaseFrom] = useState('');
  const [remark, setRemark] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async () => {
    setSubmitting(true);
    try {
      const { data: bookId } = await api.post('/books', {
        title,
        isbn,
        purchaseFrom,
        remark,
      });
      history.push(`/books/${bookId}`);
    } catch (err) {
      setError(err.message);
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <form className="form-add">
      <h2>添加新书</h2>
      <div className="form-item">
        <TextField
          id="title"
          label="标题"
          fullWidth
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
      </div>
      <div className="form-item">
        <TextField
          id="isbn"
          label="ISBN"
          fullWidth
          value={isbn}
          onChange={(e) => setIsbn(e.target.value)}
        />
      </div>
      <div className="form-item">
        <TextField
          id="purchaseFrom"
          label="购买于"
          fullWidth
          value={purchaseFrom}
          onChange={(e) => setPurchaseFrom(e.target.value)}
        />
      </div>
      <div className="form-item">
        <TextField
          multiline
          minRows={2}
          id="remark"
          fullWidth
          label="备忘"
          value={remark}
          onChange={(e) => setRemark(e.target.value)}
        />
      </div>
      {error && (
        <Alert className="form-item" severity="error">
          {error}
        </Alert>
      )}
      <div className="form-item">
        <Button
          disabled={submitting}
          onClick={handleSubmit}
          variant="contained"
          color="primary"
        >
          添加
          {submitting && (
            <CircularProgress
              style={{ color: 'white', marginLeft: '3px' }}
              size={12}
            />
          )}
        </Button>
      </div>
    </form>
  );
}
