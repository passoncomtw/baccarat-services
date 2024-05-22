import PropTypes from 'prop-types';
import { Box } from '@mui/material';

import ProfileSection from './ProfileSection';

// ==============================|| MAIN NAVBAR / HEADER ||============================== //

const Header = () => {
    return (
        <>
            <Box sx={{ flexGrow: 1 }} />
            <ProfileSection />
        </>
    );
};

Header.propTypes = {
    handleLeftDrawerToggle: PropTypes.func
};

export default Header;
