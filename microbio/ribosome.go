package microbio

// Ribosome : Here we attempt to model the Translation phase
// of gene expression done by the Ribosome
type Ribosome struct {
}

var startCodon = MakeStrand([]byte("AUG"), RIBOSE)

// Translate returns an AminoAcidStrand for a given NucleotideStrand
// Use WaitForTRNAAndGetAttachedAminoAcid(codon) to get AminoAcid
func (t *Ribosome) Translate(nucleotides NucleotideStrand) AminoAcidStrand {

	// TODO: implement
	// First, loop through nucleotides looking for START (matching startCodon)
	// then adding amino acids to an AminoAcidStrand until STOP.
	// A STOP is when WaitForTRNAAndGetAttachedAminoAcid() returns nil.
	// If reach the end of the NucleotideStrand without a STOP,
	// then consider that an error and return an empty AminoAcidStrand.
	//AminoAcidStrand := ""
	//foundStop := false
	startFound := false
	for i := 0; i < len(nucleotides.buf); i += 3 {
		nucleotides.position = i
		currentCodon := nucleotides.Codon()
		if currentCodon == nil {

			break
		}
		if currentCodon.Matches(startCodon) {

			startFound = true
			break
		}
	}
	var aminoAcidStrand AminoAcidStrand
	var emptyAminoAcidStrand AminoAcidStrand
	if startFound {
		nucleotides.position += 3
		currentCodon := nucleotides.Codon()
		currentAminoAcid := WaitForTRNAAndGetAttachedAminoAcid(currentCodon)

		for currentAminoAcid != nil {
			aminoAcidStrand = append(aminoAcidStrand, currentAminoAcid)
			nucleotides.position += 3
			currentCodon = nucleotides.Codon()
			if currentCodon != nil {
				currentAminoAcid = WaitForTRNAAndGetAttachedAminoAcid(currentCodon)
			} else {
				return emptyAminoAcidStrand
			}
		}
	}
	return aminoAcidStrand

}
