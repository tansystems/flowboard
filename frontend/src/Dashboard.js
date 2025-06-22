import * as React from 'react';
import { Card, CardContent, Typography } from '@mui/material';
import { Bar } from 'react-chartjs-2';
import { useEffect, useState } from 'react';
import { fetchUtils } from 'react-admin';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

export const Dashboard = () => {
    const [stats, setStats] = useState({ customers: 0, deals: 0, users: 0 });

    useEffect(() => {
        Promise.all([
            fetchUtils.fetchJson('http://localhost:8080/customers'),
            fetchUtils.fetchJson('http://localhost:8080/deals'),
            fetchUtils.fetchJson('http://localhost:8080/users'),
        ]).then(([customers, deals, users]) => {
            setStats({
                customers: customers.json.length,
                deals: deals.json.length,
                users: users.json.length,
            });
        });
    }, []);

    const data = {
        labels: ['Клиенты', 'Сделки', 'Пользователи'],
        datasets: [
            {
                label: 'Количество',
                data: [stats.customers, stats.deals, stats.users],
                backgroundColor: ['#1976d2', '#388e3c', '#fbc02d'],
            },
        ],
    };

    return (
        <Card>
            <CardContent>
                <Typography variant="h5" gutterBottom>
                    Добро пожаловать в CRM!
                </Typography>
                <Typography variant="body1">
                    Здесь вы можете управлять клиентами, сделками, пользователями и анализировать данные.
                </Typography>
                <div style={{ maxWidth: 600, margin: '2em auto' }}>
                    <Bar data={data} />
                </div>
            </CardContent>
        </Card>
    );
};
