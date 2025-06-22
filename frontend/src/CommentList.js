import * as React from 'react';
import { List, Datagrid, TextField, ReferenceField, EditButton, DeleteButton, TextInput } from 'react-admin';
import { isAdmin } from './helpers';

const commentFilters = [<TextInput label="Поиск по содержимому" source="q" alwaysOn key="q" />];

export const CommentList = props => (
    <List {...props} title="Комментарии" filters={commentFilters}>
        <Datagrid rowClick="edit">
            <TextField source="id" label="ID" />
            <ReferenceField source="deal_id" reference="deals" label="Сделка">
                <TextField source="title" />
            </ReferenceField>
            <ReferenceField source="user_id" reference="users" label="Пользователь">
                <TextField source="name" />
            </ReferenceField>
            <TextField source="content" label="Комментарий" />
            <EditButton />
            {isAdmin() && <DeleteButton />}
        </Datagrid>
    </List>
); 