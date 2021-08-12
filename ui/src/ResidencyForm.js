import React from 'react';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import TextField from '@material-ui/core/TextField';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Checkbox from '@material-ui/core/Checkbox';

const ResidencyForm = ({ ResidencyFormValues, changeResidencyFormValues }) => {

    return (
        <React.Fragment>
            <Typography variant="h6" gutterBottom>
                Residency Details
            </Typography>
            <Grid container spacing={3}>
                <Grid item xs={12} sm={6}>
                    <TextField
                        required
                        id="streetAddress"
                        name="streetAddress"
                        label="Street Address"
                        fullWidth
                        value={ResidencyFormValues.streetAddress}
                        onChange={(e) => changeResidencyFormValues('streetAddress', e.target.value)}
                    />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <TextField
                        required
                        id="yearsAtAddress"
                        name="yearsAtAddress"
                        type="number"
                        label="Years at address"
                        fullWidth
                        value={ResidencyFormValues.yearsAtAddress}
                        onChange={(e) => changeResidencyFormValues('yearsAtAddress', e.target.value)}
                    />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <TextField
                        required
                        id="city"
                        name="city"
                        label="City"
                        fullWidth
                        value={ResidencyFormValues.city}
                        onChange={(e) => changeResidencyFormValues('city', e.target.value)}
                    />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <TextField
                        id="state"
                        name="state"
                        label="State/Province/Region"
                        value={ResidencyFormValues.state}
                        onChange={(e) => changeResidencyFormValues('state', e.target.value)}
                        fullWidth />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <TextField
                        required
                        id="zip"
                        name="zip"
                        label="Zip / Postal code"
                        fullWidth
                        value={ResidencyFormValues.zip}
                        onChange={(e) => changeResidencyFormValues('zip', e.target.value)}
                    />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <TextField
                        required
                        id="country"
                        name="country"
                        label="Country"
                        fullWidth
                        value={ResidencyFormValues.country}
                        onChange={(e) => changeResidencyFormValues('country', e.target.value)}
                    />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <TextField
                        required
                        id="ownership"
                        name="ownership"
                        label="Ownership"
                        fullWidth
                        value={ResidencyFormValues.ownership}
                        onChange={(e) => changeResidencyFormValues('ownership', e.target.value)}
                    />
                </Grid>
                <Grid item xs={12} sm={6}>
                    <TextField
                        required
                        id="monthlyPayment"
                        name="monthlyPayment"
                        label="Monthly Payment"
                        fullWidth
                        value={ResidencyFormValues.monthlyPayment}
                        onChange={(e) => changeResidencyFormValues('monthlyPayment', e.target.value)}
                    />
                </Grid>
                <Grid item xs={12}>
                    <FormControlLabel
                        control={<Checkbox color="secondary" name="saveAddress" value="yes" />}
                        label="Use this address for payment details"
                    />
                </Grid>
            </Grid>
        </React.Fragment>
    );
}
export default ResidencyForm;