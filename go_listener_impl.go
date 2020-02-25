package example

import (
	"github.com/Applifier/golang-backend-assignment/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type GoParserListenerImpl struct {
	antlr.ParseTreeListener
	methods       []*MethodDescription
	parsingMethod *MethodDescription
	state         *State
}

func NewGoParserListenerImpl() parser.GoParserListener {
	return &GoParserListenerImpl{
		ParseTreeListener: &parser.BaseGoParserListener{},
		methods:           []*MethodDescription{},
		state:             NewState(),
	}
}

func (g *GoParserListenerImpl) EnterSourceFile(c *parser.SourceFileContext) {
}

func (g *GoParserListenerImpl) EnterPackageClause(c *parser.PackageClauseContext) {

}

func (g *GoParserListenerImpl) EnterImportDecl(c *parser.ImportDeclContext) {

}

func (g *GoParserListenerImpl) EnterImportSpec(c *parser.ImportSpecContext) {

}

func (g *GoParserListenerImpl) EnterImportPath(c *parser.ImportPathContext) {

}

func (g *GoParserListenerImpl) EnterDeclaration(c *parser.DeclarationContext) {

}

func (g *GoParserListenerImpl) EnterConstDecl(c *parser.ConstDeclContext) {

}

func (g *GoParserListenerImpl) EnterConstSpec(c *parser.ConstSpecContext) {

}

func (g *GoParserListenerImpl) EnterIdentifierList(c *parser.IdentifierListContext) {
	if g.state.IsEnable(StateEnterMethod) {
		if g.state.IsEnable(StateEnterReceiver) {
			g.parsingMethod.Receiver = c.GetText()
		} else if g.state.IsEnable(StateEnterArgument) {
			g.parsingMethod.addParameter(c.GetText())
		}
	}
}

func (g *GoParserListenerImpl) EnterExpressionList(c *parser.ExpressionListContext) {

}

func (g *GoParserListenerImpl) EnterTypeDecl(c *parser.TypeDeclContext) {
}

func (g *GoParserListenerImpl) EnterTypeSpec(c *parser.TypeSpecContext) {

}

func (g *GoParserListenerImpl) EnterFunctionDecl(c *parser.FunctionDeclContext) {
}

func (g *GoParserListenerImpl) EnterMethodDecl(c *parser.MethodDeclContext) {
	g.parsingMethod = NewMethodDescription()
	g.parsingMethod.FuncName = c.IDENTIFIER().GetText()
	g.state.Enable(StateEnterMethod)
}

func (g *GoParserListenerImpl) EnterReceiver(c *parser.ReceiverContext) {
	g.state.Enable(StateEnterReceiver)
}

func (g *GoParserListenerImpl) EnterVarDecl(c *parser.VarDeclContext) {

}

func (g *GoParserListenerImpl) EnterVarSpec(c *parser.VarSpecContext) {

}

func (g *GoParserListenerImpl) EnterBlock(c *parser.BlockContext) {

}

func (g *GoParserListenerImpl) EnterStatementList(c *parser.StatementListContext) {

}

func (g *GoParserListenerImpl) EnterStatement(c *parser.StatementContext) {

}

func (g *GoParserListenerImpl) EnterSimpleStmt(c *parser.SimpleStmtContext) {

}

func (g *GoParserListenerImpl) EnterExpressionStmt(c *parser.ExpressionStmtContext) {

}

func (g *GoParserListenerImpl) EnterSendStmt(c *parser.SendStmtContext) {

}

func (g *GoParserListenerImpl) EnterIncDecStmt(c *parser.IncDecStmtContext) {

}

func (g *GoParserListenerImpl) EnterAssignment(c *parser.AssignmentContext) {

}

func (g *GoParserListenerImpl) EnterAssign_op(c *parser.Assign_opContext) {

}

func (g *GoParserListenerImpl) EnterShortVarDecl(c *parser.ShortVarDeclContext) {

}

func (g *GoParserListenerImpl) EnterEmptyStmt(c *parser.EmptyStmtContext) {

}

func (g *GoParserListenerImpl) EnterLabeledStmt(c *parser.LabeledStmtContext) {

}

func (g *GoParserListenerImpl) EnterReturnStmt(c *parser.ReturnStmtContext) {

}

func (g *GoParserListenerImpl) EnterBreakStmt(c *parser.BreakStmtContext) {

}

func (g *GoParserListenerImpl) EnterContinueStmt(c *parser.ContinueStmtContext) {

}

func (g *GoParserListenerImpl) EnterGotoStmt(c *parser.GotoStmtContext) {

}

func (g *GoParserListenerImpl) EnterFallthroughStmt(c *parser.FallthroughStmtContext) {

}

func (g *GoParserListenerImpl) EnterDeferStmt(c *parser.DeferStmtContext) {

}

func (g *GoParserListenerImpl) EnterIfStmt(c *parser.IfStmtContext) {

}

func (g *GoParserListenerImpl) EnterSwitchStmt(c *parser.SwitchStmtContext) {

}

func (g *GoParserListenerImpl) EnterExprSwitchStmt(c *parser.ExprSwitchStmtContext) {

}

func (g *GoParserListenerImpl) EnterExprCaseClause(c *parser.ExprCaseClauseContext) {

}

func (g *GoParserListenerImpl) EnterExprSwitchCase(c *parser.ExprSwitchCaseContext) {

}

func (g *GoParserListenerImpl) EnterTypeSwitchStmt(c *parser.TypeSwitchStmtContext) {

}

func (g *GoParserListenerImpl) EnterTypeSwitchGuard(c *parser.TypeSwitchGuardContext) {

}

func (g *GoParserListenerImpl) EnterTypeCaseClause(c *parser.TypeCaseClauseContext) {

}

func (g *GoParserListenerImpl) EnterTypeSwitchCase(c *parser.TypeSwitchCaseContext) {

}

func (g *GoParserListenerImpl) EnterTypeList(c *parser.TypeListContext) {

}

func (g *GoParserListenerImpl) EnterSelectStmt(c *parser.SelectStmtContext) {

}

func (g *GoParserListenerImpl) EnterCommClause(c *parser.CommClauseContext) {

}

func (g *GoParserListenerImpl) EnterCommCase(c *parser.CommCaseContext) {

}

func (g *GoParserListenerImpl) EnterRecvStmt(c *parser.RecvStmtContext) {

}

func (g *GoParserListenerImpl) EnterForStmt(c *parser.ForStmtContext) {

}

func (g *GoParserListenerImpl) EnterForClause(c *parser.ForClauseContext) {

}

func (g *GoParserListenerImpl) EnterRangeClause(c *parser.RangeClauseContext) {

}

func (g *GoParserListenerImpl) EnterGoStmt(c *parser.GoStmtContext) {

}

func (g *GoParserListenerImpl) EnterType_(c *parser.Type_Context) {
	if g.state.IsEnable(StateEnterMethod) {
		if g.state.IsEnable(StateEnterReceiver) {
			g.parsingMethod.ReceiverType = c.GetText()
		} else if g.state.IsEnable(StateEnterSignature) && g.state.IsEnable(StateEnterArgument) {
			g.parsingMethod.addType(c.GetText())
		}
	}
}

func (g *GoParserListenerImpl) EnterTypeName(c *parser.TypeNameContext) {

}

func (g *GoParserListenerImpl) EnterTypeLit(c *parser.TypeLitContext) {

}

func (g *GoParserListenerImpl) EnterArrayType(c *parser.ArrayTypeContext) {

}

func (g *GoParserListenerImpl) EnterArrayLength(c *parser.ArrayLengthContext) {

}

func (g *GoParserListenerImpl) EnterElementType(c *parser.ElementTypeContext) {

}

func (g *GoParserListenerImpl) EnterPointerType(c *parser.PointerTypeContext) {

}

func (g *GoParserListenerImpl) EnterInterfaceType(c *parser.InterfaceTypeContext) {

}

func (g *GoParserListenerImpl) EnterSliceType(c *parser.SliceTypeContext) {

}

func (g *GoParserListenerImpl) EnterMapType(c *parser.MapTypeContext) {

}

func (g *GoParserListenerImpl) EnterChannelType(c *parser.ChannelTypeContext) {

}

func (g *GoParserListenerImpl) EnterMethodSpec(c *parser.MethodSpecContext) {

}

func (g *GoParserListenerImpl) EnterFunctionType(c *parser.FunctionTypeContext) {

}

func (g *GoParserListenerImpl) EnterSignature(c *parser.SignatureContext) {
	if g.state.IsEnable(StateEnterMethod) {
		g.state.Enable(StateEnterSignature)
	}
}

func (g *GoParserListenerImpl) EnterResult(c *parser.ResultContext) {

}

func (g *GoParserListenerImpl) EnterParameters(c *parser.ParametersContext) {
}

func (g *GoParserListenerImpl) EnterParameterDecl(c *parser.ParameterDeclContext) {
	if g.state.IsEnable(StateEnterMethod) {
		g.state.Enable(StateEnterArgument)
	}
}

func (g *GoParserListenerImpl) EnterExpression(c *parser.ExpressionContext) {

}

func (g *GoParserListenerImpl) EnterPrimaryExpr(c *parser.PrimaryExprContext) {

}

func (g *GoParserListenerImpl) EnterUnaryExpr(c *parser.UnaryExprContext) {

}

func (g *GoParserListenerImpl) EnterConversion(c *parser.ConversionContext) {

}

func (g *GoParserListenerImpl) EnterOperand(c *parser.OperandContext) {

}

func (g *GoParserListenerImpl) EnterLiteral(c *parser.LiteralContext) {

}

func (g *GoParserListenerImpl) EnterBasicLit(c *parser.BasicLitContext) {

}

func (g *GoParserListenerImpl) EnterInteger(c *parser.IntegerContext) {

}

func (g *GoParserListenerImpl) EnterOperandName(c *parser.OperandNameContext) {

}

func (g *GoParserListenerImpl) EnterQualifiedIdent(c *parser.QualifiedIdentContext) {

}

func (g *GoParserListenerImpl) EnterCompositeLit(c *parser.CompositeLitContext) {

}

func (g *GoParserListenerImpl) EnterLiteralType(c *parser.LiteralTypeContext) {

}

func (g *GoParserListenerImpl) EnterLiteralValue(c *parser.LiteralValueContext) {

}

func (g *GoParserListenerImpl) EnterElementList(c *parser.ElementListContext) {

}

func (g *GoParserListenerImpl) EnterKeyedElement(c *parser.KeyedElementContext) {

}

func (g *GoParserListenerImpl) EnterKey(c *parser.KeyContext) {

}

func (g *GoParserListenerImpl) EnterElement(c *parser.ElementContext) {

}

func (g *GoParserListenerImpl) EnterStructType(c *parser.StructTypeContext) {
}

func (g *GoParserListenerImpl) EnterFieldDecl(c *parser.FieldDeclContext) {

}

func (g *GoParserListenerImpl) EnterString_(c *parser.String_Context) {

}

func (g *GoParserListenerImpl) EnterAnonymousField(c *parser.AnonymousFieldContext) {

}

func (g *GoParserListenerImpl) EnterFunctionLit(c *parser.FunctionLitContext) {

}

func (g *GoParserListenerImpl) EnterIndex(c *parser.IndexContext) {

}

func (g *GoParserListenerImpl) EnterSlice(c *parser.SliceContext) {

}

func (g *GoParserListenerImpl) EnterTypeAssertion(c *parser.TypeAssertionContext) {

}

func (g *GoParserListenerImpl) EnterArguments(c *parser.ArgumentsContext) {

}

func (g *GoParserListenerImpl) EnterMethodExpr(c *parser.MethodExprContext) {
}

func (g *GoParserListenerImpl) EnterReceiverType(c *parser.ReceiverTypeContext) {
}

func (g *GoParserListenerImpl) EnterEos(c *parser.EosContext) {

}

func (g *GoParserListenerImpl) ExitSourceFile(c *parser.SourceFileContext) {

}

func (g *GoParserListenerImpl) ExitPackageClause(c *parser.PackageClauseContext) {

}

func (g *GoParserListenerImpl) ExitImportDecl(c *parser.ImportDeclContext) {

}

func (g *GoParserListenerImpl) ExitImportSpec(c *parser.ImportSpecContext) {

}

func (g *GoParserListenerImpl) ExitImportPath(c *parser.ImportPathContext) {

}

func (g *GoParserListenerImpl) ExitDeclaration(c *parser.DeclarationContext) {

}

func (g *GoParserListenerImpl) ExitConstDecl(c *parser.ConstDeclContext) {

}

func (g *GoParserListenerImpl) ExitConstSpec(c *parser.ConstSpecContext) {

}

func (g *GoParserListenerImpl) ExitIdentifierList(c *parser.IdentifierListContext) {

}

func (g *GoParserListenerImpl) ExitExpressionList(c *parser.ExpressionListContext) {

}

func (g *GoParserListenerImpl) ExitTypeDecl(c *parser.TypeDeclContext) {

}

func (g *GoParserListenerImpl) ExitTypeSpec(c *parser.TypeSpecContext) {

}

func (g *GoParserListenerImpl) ExitFunctionDecl(c *parser.FunctionDeclContext) {

}

func (g *GoParserListenerImpl) ExitMethodDecl(c *parser.MethodDeclContext) {
	g.methods = append(g.methods, g.parsingMethod)
	g.parsingMethod = nil
	g.state = NewState()
}

func (g *GoParserListenerImpl) ExitReceiver(c *parser.ReceiverContext) {
	g.state.Disable(StateEnterReceiver)
}

func (g *GoParserListenerImpl) ExitVarDecl(c *parser.VarDeclContext) {

}

func (g *GoParserListenerImpl) ExitVarSpec(c *parser.VarSpecContext) {

}

func (g *GoParserListenerImpl) ExitBlock(c *parser.BlockContext) {

}

func (g *GoParserListenerImpl) ExitStatementList(c *parser.StatementListContext) {

}

func (g *GoParserListenerImpl) ExitStatement(c *parser.StatementContext) {

}

func (g *GoParserListenerImpl) ExitSimpleStmt(c *parser.SimpleStmtContext) {

}

func (g *GoParserListenerImpl) ExitExpressionStmt(c *parser.ExpressionStmtContext) {

}

func (g *GoParserListenerImpl) ExitSendStmt(c *parser.SendStmtContext) {

}

func (g *GoParserListenerImpl) ExitIncDecStmt(c *parser.IncDecStmtContext) {

}

func (g *GoParserListenerImpl) ExitAssignment(c *parser.AssignmentContext) {

}

func (g *GoParserListenerImpl) ExitAssign_op(c *parser.Assign_opContext) {

}

func (g *GoParserListenerImpl) ExitShortVarDecl(c *parser.ShortVarDeclContext) {

}

func (g *GoParserListenerImpl) ExitEmptyStmt(c *parser.EmptyStmtContext) {

}

func (g *GoParserListenerImpl) ExitLabeledStmt(c *parser.LabeledStmtContext) {

}

func (g *GoParserListenerImpl) ExitReturnStmt(c *parser.ReturnStmtContext) {

}

func (g *GoParserListenerImpl) ExitBreakStmt(c *parser.BreakStmtContext) {

}

func (g *GoParserListenerImpl) ExitContinueStmt(c *parser.ContinueStmtContext) {

}

func (g *GoParserListenerImpl) ExitGotoStmt(c *parser.GotoStmtContext) {

}

func (g *GoParserListenerImpl) ExitFallthroughStmt(c *parser.FallthroughStmtContext) {

}

func (g *GoParserListenerImpl) ExitDeferStmt(c *parser.DeferStmtContext) {

}

func (g *GoParserListenerImpl) ExitIfStmt(c *parser.IfStmtContext) {

}

func (g *GoParserListenerImpl) ExitSwitchStmt(c *parser.SwitchStmtContext) {

}

func (g *GoParserListenerImpl) ExitExprSwitchStmt(c *parser.ExprSwitchStmtContext) {

}

func (g *GoParserListenerImpl) ExitExprCaseClause(c *parser.ExprCaseClauseContext) {

}

func (g *GoParserListenerImpl) ExitExprSwitchCase(c *parser.ExprSwitchCaseContext) {

}

func (g *GoParserListenerImpl) ExitTypeSwitchStmt(c *parser.TypeSwitchStmtContext) {

}

func (g *GoParserListenerImpl) ExitTypeSwitchGuard(c *parser.TypeSwitchGuardContext) {

}

func (g *GoParserListenerImpl) ExitTypeCaseClause(c *parser.TypeCaseClauseContext) {

}

func (g *GoParserListenerImpl) ExitTypeSwitchCase(c *parser.TypeSwitchCaseContext) {

}

func (g *GoParserListenerImpl) ExitTypeList(c *parser.TypeListContext) {

}

func (g *GoParserListenerImpl) ExitSelectStmt(c *parser.SelectStmtContext) {

}

func (g *GoParserListenerImpl) ExitCommClause(c *parser.CommClauseContext) {

}

func (g *GoParserListenerImpl) ExitCommCase(c *parser.CommCaseContext) {

}

func (g *GoParserListenerImpl) ExitRecvStmt(c *parser.RecvStmtContext) {

}

func (g *GoParserListenerImpl) ExitForStmt(c *parser.ForStmtContext) {

}

func (g *GoParserListenerImpl) ExitForClause(c *parser.ForClauseContext) {

}

func (g *GoParserListenerImpl) ExitRangeClause(c *parser.RangeClauseContext) {

}

func (g *GoParserListenerImpl) ExitGoStmt(c *parser.GoStmtContext) {

}

func (g *GoParserListenerImpl) ExitType_(c *parser.Type_Context) {

}

func (g *GoParserListenerImpl) ExitTypeName(c *parser.TypeNameContext) {

}

func (g *GoParserListenerImpl) ExitTypeLit(c *parser.TypeLitContext) {

}

func (g *GoParserListenerImpl) ExitArrayType(c *parser.ArrayTypeContext) {

}

func (g *GoParserListenerImpl) ExitArrayLength(c *parser.ArrayLengthContext) {

}

func (g *GoParserListenerImpl) ExitElementType(c *parser.ElementTypeContext) {

}

func (g *GoParserListenerImpl) ExitPointerType(c *parser.PointerTypeContext) {

}

func (g *GoParserListenerImpl) ExitInterfaceType(c *parser.InterfaceTypeContext) {

}

func (g *GoParserListenerImpl) ExitSliceType(c *parser.SliceTypeContext) {

}

func (g *GoParserListenerImpl) ExitMapType(c *parser.MapTypeContext) {

}

func (g *GoParserListenerImpl) ExitChannelType(c *parser.ChannelTypeContext) {

}

func (g *GoParserListenerImpl) ExitMethodSpec(c *parser.MethodSpecContext) {

}

func (g *GoParserListenerImpl) ExitFunctionType(c *parser.FunctionTypeContext) {

}

func (g *GoParserListenerImpl) ExitSignature(c *parser.SignatureContext) {
}

func (g *GoParserListenerImpl) ExitResult(c *parser.ResultContext) {

}

func (g *GoParserListenerImpl) ExitParameters(c *parser.ParametersContext) {

}

func (g *GoParserListenerImpl) ExitParameterDecl(c *parser.ParameterDeclContext) {
	if g.state.IsEnable(StateEnterMethod) {
		g.state.Disable(StateEnterArgument)
	}
}

func (g *GoParserListenerImpl) ExitExpression(c *parser.ExpressionContext) {

}

func (g *GoParserListenerImpl) ExitPrimaryExpr(c *parser.PrimaryExprContext) {

}

func (g *GoParserListenerImpl) ExitUnaryExpr(c *parser.UnaryExprContext) {

}

func (g *GoParserListenerImpl) ExitConversion(c *parser.ConversionContext) {

}

func (g *GoParserListenerImpl) ExitOperand(c *parser.OperandContext) {

}

func (g *GoParserListenerImpl) ExitLiteral(c *parser.LiteralContext) {

}

func (g *GoParserListenerImpl) ExitBasicLit(c *parser.BasicLitContext) {

}

func (g *GoParserListenerImpl) ExitInteger(c *parser.IntegerContext) {

}

func (g *GoParserListenerImpl) ExitOperandName(c *parser.OperandNameContext) {

}

func (g *GoParserListenerImpl) ExitQualifiedIdent(c *parser.QualifiedIdentContext) {

}

func (g *GoParserListenerImpl) ExitCompositeLit(c *parser.CompositeLitContext) {

}

func (g *GoParserListenerImpl) ExitLiteralType(c *parser.LiteralTypeContext) {

}

func (g *GoParserListenerImpl) ExitLiteralValue(c *parser.LiteralValueContext) {

}

func (g *GoParserListenerImpl) ExitElementList(c *parser.ElementListContext) {

}

func (g *GoParserListenerImpl) ExitKeyedElement(c *parser.KeyedElementContext) {

}

func (g *GoParserListenerImpl) ExitKey(c *parser.KeyContext) {

}

func (g *GoParserListenerImpl) ExitElement(c *parser.ElementContext) {

}

func (g *GoParserListenerImpl) ExitStructType(c *parser.StructTypeContext) {

}

func (g *GoParserListenerImpl) ExitFieldDecl(c *parser.FieldDeclContext) {

}

func (g *GoParserListenerImpl) ExitString_(c *parser.String_Context) {

}

func (g *GoParserListenerImpl) ExitAnonymousField(c *parser.AnonymousFieldContext) {

}

func (g *GoParserListenerImpl) ExitFunctionLit(c *parser.FunctionLitContext) {

}

func (g *GoParserListenerImpl) ExitIndex(c *parser.IndexContext) {

}

func (g *GoParserListenerImpl) ExitSlice(c *parser.SliceContext) {

}

func (g *GoParserListenerImpl) ExitTypeAssertion(c *parser.TypeAssertionContext) {

}

func (g *GoParserListenerImpl) ExitArguments(c *parser.ArgumentsContext) {
}

func (g *GoParserListenerImpl) ExitMethodExpr(c *parser.MethodExprContext) {

}

func (g *GoParserListenerImpl) ExitReceiverType(c *parser.ReceiverTypeContext) {

}

func (g *GoParserListenerImpl) ExitEos(c *parser.EosContext) {

}
