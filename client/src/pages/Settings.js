import { Helmet } from 'react-helmet';
import { Box, Container } from '@material-ui/core';
import SettingsPassword from '../components/settings/SettingsPassword';

const SettingsView = () => (
  <>
    <Helmet>
      <title>Configurações | Amarlinda</title>
    </Helmet>
    <Box
      sx={{
        backgroundColor: 'background.default',
        minHeight: '100%',
        py: 3
      }}
    >
      <Container maxWidth="lg">
        <Box sx={{ pt: 3 }}>
          <SettingsPassword />
        </Box>
      </Container>
    </Box>
  </>
);

export default SettingsView;
