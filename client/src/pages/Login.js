import { useNavigate } from 'react-router-dom';
import { Helmet } from 'react-helmet';
import * as Yup from 'yup';
import { Formik } from 'formik';
import Snackbar from '@material-ui/core/Snackbar';
import IconButton from '@material-ui/core/IconButton';
import CloseIcon from '@material-ui/icons/Close';
import {
  Box,
  Button,
  Container,
  TextField,
  Typography
} from '@material-ui/core';

const Login = () => {
  const navigate = useNavigate();

  return (
    <>
      <Helmet>
        <title>Login</title>
      </Helmet>
      <Box
        sx={{
          backgroundColor: 'background.default',
          display: 'flex',
          flexDirection: 'column',
          height: '100%',
          justifyContent: 'center'
        }}
      >
        <Container maxWidth="sm">
          <Formik
            initialValues={{
              email: 'login@amarlinda.com.br',
              password: 'Password123',
              error: 'asdfsdfsd'
            }}
            validationSchema={Yup.object().shape({
              email: Yup.string().email('O e-mail precisa ser válido').max(255).required('E-mail é obrigatório'),
              password: Yup.string().max(255).required('Senha é obrigatório')
            })}
            onSubmit={async (values) => {
              const { email, password } = values;

              const response = await fetch('http://localhost:8000/api/login', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
                body: JSON.stringify({
                  email,
                  password
                })
              });

              if (response.status !== 200) {
                const content = await response.json();
                console.log(content.message);
              }

              navigate('/app/dashboard', { replace: true });
            }}
          >
            {({
              errors,
              handleBlur,
              handleChange,
              handleSubmit,
              isSubmitting,
              touched,
              values
            }) => (
              <form onSubmit={handleSubmit}>
                <Box sx={{ mb: 3 }}>
                  <Typography
                    color="textPrimary"
                    variant="h2"
                    align="center"
                  >
                    Acesso a loja Amarlinda
                  </Typography>
                  <Typography
                    color="textSecondary"
                    gutterBottom
                    variant="body2"
                    align="center"
                  >
                    Acesse o sistema informando seu e-mail e senha.
                  </Typography>
                </Box>
                <TextField
                  error={Boolean(touched.email && errors.email)}
                  fullWidth
                  helperText={touched.email && errors.email}
                  label="E-mail"
                  margin="normal"
                  name="email"
                  onBlur={handleBlur}
                  onChange={handleChange}
                  type="email"
                  value={values.email}
                  variant="outlined"
                />
                <TextField
                  error={Boolean(touched.password && errors.password)}
                  fullWidth
                  helperText={touched.password && errors.password}
                  label="Senha"
                  margin="normal"
                  name="password"
                  onBlur={handleBlur}
                  onChange={handleChange}
                  type="password"
                  value={values.password}
                  variant="outlined"
                />
                {values.error ? (
                  <Snackbar
                    anchorOrigin={{
                      vertical: 'bottom',
                      horizontal: 'left',
                    }}
                    open
                    autoHideDuration={6000}
                    ContentProps={{
                      'aria-describedby': 'message-id',
                    }}
                    message={<span id="message-id">{values.error}</span>}
                    action={[
                      <Button key="undo" color="secondary" size="small">
                        UNDO
                      </Button>,
                      <IconButton
                        key="close"
                        aria-label="Close"
                        color="inherit"
                      >
                        <CloseIcon />
                      </IconButton>,
                    ]}
                  />
                ) : null}
                <Box sx={{ py: 2 }}>
                  <Button
                    color="primary"
                    disabled={isSubmitting}
                    fullWidth
                    size="large"
                    type="submit"
                    variant="contained"
                  >
                    Entrar
                  </Button>
                </Box>
              </form>
            )}
          </Formik>
        </Container>
      </Box>
    </>
  );
};

export default Login;
