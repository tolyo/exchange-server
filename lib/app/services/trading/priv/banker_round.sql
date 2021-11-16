-- round-half-even aka Banker's rounding
CREATE OR REPLACE FUNCTION 
    banker_round(val NUMERIC, prec INTEGER)
    returns NUMERIC

LANGUAGE 'plpgsql' IMMUTABLE
AS $$
DECLARE
    retval NUMERIC;
    difference NUMERIC;
    even BOOLEAN;
BEGIN
    IF val IS NULL THEN
        RAISE EXCEPTION 'val_cannot_be_null';
    END IF; 

    IF prec IS NULL THEN
        RAISE EXCEPTION 'prec_cannot_be_null';
    END IF; 

    retval := round(val,prec);
    difference := retval-val;
    IF abs(difference)*(10::NUMERIC^prec) = 0.5::NUMERIC THEN
        even := (retval * (10::NUMERIC^prec)) % 2::NUMERIC = 0::NUMERIC;
        IF NOT even THEN
            retval := round(val-difference,prec);
        END IF;
    END IF;
    RETURN retval;
END;
$$;