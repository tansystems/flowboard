import * as React from 'react';
import { Edit, SimpleForm, TextInput, SelectInput } from 'react-admin';

export const UserEdit = props => (
    <Edit {...props} title="Редактировать пользователя">
        <SimpleForm>
            <TextInput disabled source="id" label="ID" />
            <TextInput source="name" label="Имя" />
            <TextInput source="email" label="Email" />
            <SelectInput source="role" label="Роль" choices={[
                { id: 'user', name: 'Пользователь' },
                { id: 'admin', name: 'Администратор' }
            ]} />
        </SimpleForm>
    </Edit>
); 