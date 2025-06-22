import * as React from 'react';
import { Create, SimpleForm, TextInput } from 'react-admin';

export const TagCreate = props => (
    <Create {...props} title="Создать тег">
        <SimpleForm>
            <TextInput source="name" label="Название" />
        </SimpleForm>
    </Create>
); 