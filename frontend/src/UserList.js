import * as React from 'react';
import { List, Datagrid, TextField, EmailField, EditButton, DeleteButton, TextInput } from 'react-admin';
import { isAdmin } from './helpers';

const userFilters = [<TextInput label="Поиск по email" source="q" alwaysOn key="q" />];

export const UserList = props => (
    <List {...props} title="Пользователи" filters={userFilters}>
        <Datagrid rowClick="edit">
            <TextField source="id" label="ID" />
            <TextField source="name" label="Имя" />
            <EmailField source="email" label="Email" />
            <TextField source="role" label="Роль" />
            <EditButton />
            {isAdmin() && <DeleteButton />}
        </Datagrid>
    </List>
); 