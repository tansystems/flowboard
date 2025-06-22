import * as React from 'react';
import { List, Datagrid, TextField, EmailField, EditButton, DeleteButton } from 'react-admin';

export const CustomerList = (props) => (
    <List {...props} title="Клиенты">
        <Datagrid rowClick="edit">
            <TextField source="id" label="ID" />
            <TextField source="name" label="Имя" />
            <EmailField source="email" label="Email" />
            <TextField source="phone" label="Телефон" />
            <TextField source="company" label="Компания" />
            <EditButton />
            <DeleteButton />
        </Datagrid>
    </List>
); 