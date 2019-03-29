package model


type SMContext struct {
	// TODO
}

type SMFPDUSessions struct{
	SMContexts map[string]SMContext
	// TODO
}

func (sm *SMFPDUSessions)CreatePduSession(reqBody *SmContextCreateBody)error{
	return nil
}