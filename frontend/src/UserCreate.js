import * as React from 'react';
import { Create, SimpleForm, TextInput, SelectInput } from 'react-admin';

export const UserCreate = props => (
    <Create {...props} title="Создать пользователя">
        <SimpleForm>
            <TextInput source="name" label="Имя" />
            <TextInput source="email" label="Email" />
            <TextInput source="passwordHash" label="Пароль" type="password" />
            <SelectInput source="role" label="Роль" choices={[
                { id: 'user', name: 'Пользователь' },
                { id: 'admin', name: 'Администратор' }
            ]} />
        </SimpleForm>
    </Create>
); 