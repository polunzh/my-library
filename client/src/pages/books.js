import React from 'react';
import { useHistory } from 'react-router-dom';
import { useAsync } from 'react-use';
import { createTheme, MuiThemeProvider } from '@material-ui/core/styles';
import CircularProgress from '@material-ui/core/CircularProgress';
import Alert from '@material-ui/lab/Alert';
import { Button } from '@material-ui/core';
import MUIDataTable from 'mui-datatables';

import api from '../libs/api';
import { formatDatetime } from '../libs/util';

export default function Books() {
  const history = useHistory();
  const booksState = useAsync(async () => {
    const resp = await api.get('/books');
    return resp.data;
  });

  if (booksState.loading) {
    return <CircularProgress />;
  }

  if (booksState.error) {
    return <Alert severity="error">{booksState.error.message}</Alert>;
  }

  const columns = [
    {
      name: 'Isbn',
      label: 'ISBN',
      options: {
        filter: true,
        sort: false,
      },
    },
    {
      name: 'Title',
      label: '标题',
      options: {
        filter: true,
        sort: false,
      },
    },
    {
      name: 'PurchaseFrom',
      label: '购买于',
      options: {
        filter: true,
        sort: false,
      },
    },
    {
      name: 'Remark',
      label: '备注',
      options: {
        filter: false,
        sort: false,
      },
    },
    {
      name: 'CreatedAt',
      label: '添加时间',
      options: {
        filter: true,
        sort: true,
        customBodyRender: (value) => formatDatetime(value),
      },
    },
    {
      name: 'UpdatedAt',
      label: '更新时间',
      options: {
        filter: true,
        sort: true,
        customBodyRender: (value) => formatDatetime(value),
      },
    },
  ];

  const options = {
    download: false,
    print: false,
    selectableRowsHideCheckboxes: true,
    onRowClick: (_, { dataIndex }) => {
      const { Id } = booksState.value[dataIndex];
      history.push(`/books/${Id}`);
    },
    textLabels: {
      body: {
        noMatch: '还没有添加书籍',
      },
    },
  };

  const theme = createTheme({
    overrides: {
      MUIDataTableBodyRow: {
        root: {
          cursor: 'pointer',
        },
      },
    },
  });

  return (
    <>
      <div className="add-book">
        <Button variant="contained" href="/add" color="primary">
          添加新书
        </Button>
      </div>
      <MuiThemeProvider theme={theme}>
        <MUIDataTable
          className="book-list"
          title={'所有书目'}
          data={booksState.value}
          columns={columns}
          options={options}
        />
      </MuiThemeProvider>
    </>
  );
}
