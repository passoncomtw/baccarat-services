import { lazy } from 'react';

// project imports
import MainLayout from 'layout/MainLayout';
import Loadable from 'ui-component/Loadable';
import ProtectedRoute from 'components/ProtectedRoute';

// dashboard routing
const PokerScreen = Loadable(lazy(() => import('screens/PokerScreen')));
const TableScreen = Loadable(lazy(() => import('screens/TableScreen')));
const RoomScreen = Loadable(lazy(() => import('screens/RoomScreen')));
const NotFoundScreen = Loadable(lazy(() => import('screens/NotFoundScreen')));

// ==============================|| MAIN ROUTING ||============================== //

const MainRoutes = {
    path: '/',
    element: <MainLayout />,
    children: [
        {
            path: '/',
            element: <ProtectedRoute children={<RoomScreen />} />
        },
        {
            path: '/poker',
            element: <ProtectedRoute children={<PokerScreen />} />
        },
        {
            path: '/table',
            element: <ProtectedRoute children={<TableScreen />} />
        },
        {
            path: '/*',
            element: <ProtectedRoute children={<NotFoundScreen />} />
        },
    ]
};

export default MainRoutes;
