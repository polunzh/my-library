import React from 'react';
import { useParams } from 'react-router-dom';
import useAsync from 'react-use/lib/useAsync';
import CircularProgress from '@material-ui/core/CircularProgress';
import Alert from '@material-ui/lab/Alert';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';

import api from '../libs/api';
import { formatDatetime } from '../libs/util';

export default function Book() {
  const { id } = useParams();
  const bookState = useAsync(async () => {
    const { data } = await api.get(`/books/${id}`);
    return data;
  });

  if (bookState.loading) {
    return <CircularProgress />;
  }

  if (bookState.error) {
    return <Alert severity="error">{bookState.error.message}</Alert>;
  }

  const { value: book } = bookState;
  const bookItems = [
    { name: '标题', value: book.Title },
    { name: 'ISBN', value: book.Isbn },
    { name: '购买自哪里', value: book.PurchaseFrom },
    { name: '备忘', value: book.Remark },
    { name: '创建时间', value: formatDatetime(book.CreatedAt) },
    { name: '更新时间', value: formatDatetime(book.UpdatedAt) },
  ];

  return (
    <List>
      {bookItems.map((item) => (
        <ListItem
          key={item.name}
          style={{
            fontSize: '1rem',
            fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
            fontWeight: '400',
            lineHeight: '1.5',
            letterSpacing: '0.00938em',
          }}
        >
          <label style={{ minWidth: '150px' }}>{item.name}</label>
          <span style={{ minWidth: '150px' }}>{item.value || '-'}</span>
        </ListItem>
      ))}
    </List>
  );
}
