import * as React from 'react';
import { Create, SimpleForm, TextInput, ReferenceInput, SelectInput } from 'react-admin';

export const CommentCreate = props => (
    <Create {...props} title="Создать комментарий">
        <SimpleForm>
            <ReferenceInput source="deal_id" reference="deals" label="Сделка">
                <SelectInput optionText="title" />
            </ReferenceInput>
            <ReferenceInput source="user_id" reference="users" label="Пользователь">
                <SelectInput optionText="name" />
            </ReferenceInput>
            <TextInput source="content" label="Комментарий" />
        </SimpleForm>
    </Create>
); 